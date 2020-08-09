package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/mrdulin/graphql-go-cnode/schema"
)

const (
	gqlpath = "/graphql"
	port    = 3000
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func RootObjectFn(ctx context.Context, r *http.Request) map[string]interface{} {
	auth := r.Header.Get("authorization")
	return map[string]interface{}{
		"auth": auth,
	}
}

func main() {
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: schema.RootQuery,
	})

	h := handler.New(&handler.Config{
		Schema:       &schema,
		Pretty:       true,
		GraphiQL:     false,
		Playground:   true,
		RootObjectFn: RootObjectFn,
	})

	http.Handle(gqlpath, h)
	fmt.Printf("Access the web app via browser at http://localhost:%d%s", port, gqlpath)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
