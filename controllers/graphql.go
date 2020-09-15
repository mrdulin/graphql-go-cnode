package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/mrdulin/graphql-go-cnode/services"
)

type GraphqlHandlerOptions struct {
	Query        *graphql.Object
	Mutation     *graphql.Object
	Services     *services.Container
	RootObjectFn handler.RootObjectFn
}

func NewGraphqlHandler(options GraphqlHandlerOptions) *handler.Handler {
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    options.Query,
		Mutation: options.Mutation,
	})

	if err != nil {
		log.Fatalln(err)
	}

	h := handler.New(&handler.Config{
		Schema:       &graphqlSchema,
		Pretty:       true,
		GraphiQL:     false,
		Playground:   true,
		RootObjectFn: options.RootObjectFn,
	})

	return h
}
func NewRootObjectFn(services *services.Container) handler.RootObjectFn {
	return func(ctx context.Context, r *http.Request) map[string]interface{} {
		auth := r.Header.Get("authorization")
		return map[string]interface{}{
			"auth":     auth,
			"services": services,
		}
	}
}
