package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
)

type ResponseData struct {
	Data interface{} `json:"data"`
}
type ResponseStatus struct {
	Success bool `json:"success"`
}

type RequestGetter func(url string) (interface{}, error)
type RequestPoster func(url string, body interface{}) (interface{}, error)

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

// RequestGet send GET HTTP request
func RequestGet(url string) (interface{}, error) {
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
func RequestPost(url string, body interface{}) (interface{}, error) {
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
	var data interface{}
	json.Unmarshal(respBody, &data)
	return data, nil
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
