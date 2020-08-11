package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/graphql-go/graphql"
)

type ResponseData struct {
	Data interface{} `json:"data"`
}
type ResponseStatus struct {
	Success bool `json:"success"`
}

type IHttpClient interface {
	Get(url string) (interface{}, error)
	Post(url string, body interface{}) ([]byte, error)
}

type HttpClient struct {
	IHttpClient
}

// RequestGet send GET HTTP request
func (h *HttpClient) Get(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data Response
	json.Unmarshal(body, &data)
	return data, nil
}

// RequestPost send POST HTTP request
func (h *HttpClient) Post(url string, body interface{}) ([]byte, error) {
	jsonValue, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//var data interface{}
	//json.Unmarshal(respBody, &data)
	//if err != nil {
	//	return nil, err
	//}
	return respBody, nil
}

// Response cnode API response struct
type Response struct {
	ResponseStatus
	ResponseData
}

// PrintPretty print pretty
func PrintPretty(x interface{}) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}

// MergeMap merge two maps without mutating anyone
func MergeGraphqlFields(a map[string]*graphql.Field, b map[string]*graphql.Field) graphql.Fields {
	m := map[string]*graphql.Field{}
	for k, v := range b {
		m[k] = v
	}
	for k, v := range a {
		m[k] = v
	}
	return m
}

func mergeGraphqlFields(src map[string]*graphql.Field, dest map[string]*graphql.Field, c chan map[string]*graphql.Field, mu *sync.Mutex) {
	for k, v := range src {
		mu.Lock()
		dest[k] = v
		mu.Unlock()
	}
	c <- dest
}

func MergeGraphqlFieldsConcurrency(a map[string]*graphql.Field, b map[string]*graphql.Field) graphql.Fields {
	var mutex sync.Mutex
	m := map[string]*graphql.Field{}
	c := make(chan map[string]*graphql.Field)
	go mergeGraphqlFields(a, m, c, &mutex)
	go mergeGraphqlFields(b, m, c, &mutex)
	_, _ = <-c, <-c
	return m
}

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
