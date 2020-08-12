package utils

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/graphql-go/graphql"
)

// PrintPretty print pretty
// For debug purpose
func PrintPretty(x interface{}) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}

// MergeMap merge two maps without mutating anyone
func MergeGraphqlFields(s ...map[string]*graphql.Field) graphql.Fields {
	m := map[string]*graphql.Field{}
	for _, mm := range s {
		for k, v := range mm {
			m[k] = v
		}
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
