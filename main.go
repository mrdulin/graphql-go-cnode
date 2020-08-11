package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/joho/godotenv"

	svcs "github.com/mrdulin/graphql-go-cnode/services"
	"github.com/mrdulin/graphql-go-cnode/utils"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/mrdulin/graphql-go-cnode/schema"
)

var (
	services   *svcs.Container
	envFetcher utils.EnvFetcher
	gqlpath    = "/graphql"
	port       = 3000
)

func init() {
	var err error
	envFetcher = utils.NewOsEnvFetcher(godotenv.Load)
	apiBaseUrl := envFetcher.Getenv("API_BASE_URL")
	httpClient := utils.HttpClient{}
	userService := svcs.NewUserService(&httpClient, apiBaseUrl)
	topicService := svcs.NewTopicService(&httpClient, apiBaseUrl)
	gqlpath = envFetcher.Getenv("GRAPHQL_PATH")
	port, err = strconv.Atoi(envFetcher.Getenv("PORT"))
	if err != nil {
		log.Fatalln("Convert PORT string to int error", err)
	}
	services = &svcs.Container{UserService: userService, TopicService: topicService}
}

func RootObjectFn(ctx context.Context, r *http.Request) map[string]interface{} {
	auth := r.Header.Get("authorization")
	return map[string]interface{}{
		"auth":     auth,
		"services": services,
	}
}

func main() {

	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: schema.RootQuery,
	})

	if err != nil {
		log.Fatalln(err)
	}

	h := handler.New(&handler.Config{
		Schema:       &graphqlSchema,
		Pretty:       true,
		GraphiQL:     false,
		Playground:   true,
		RootObjectFn: RootObjectFn,
	})

	http.Handle(gqlpath, h)
	fmt.Printf("Access the web app via browser at http://localhost:%d%s", port, gqlpath)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
