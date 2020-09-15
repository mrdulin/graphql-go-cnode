package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/mrdulin/graphql-go-cnode/controllers"
	"github.com/mrdulin/graphql-go-cnode/schema"
	svcs "github.com/mrdulin/graphql-go-cnode/services"
	"github.com/mrdulin/graphql-go-cnode/utils"
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
	messageService := svcs.NewMessageService(&httpClient, apiBaseUrl)
	gqlpath = envFetcher.Getenv("GRAPHQL_PATH")
	port, err = strconv.Atoi(envFetcher.Getenv("PORT"))
	if err != nil {
		log.Fatalln("Convert PORT string to int error", err)
	}
	services = &svcs.Container{
		UserService:    userService,
		TopicService:   topicService,
		MessageService: messageService,
	}
}

func main() {

	rootObjectFn := controllers.NewRootObjectFn(services)
	graphqlHandler := controllers.NewGraphqlHandler(
		controllers.GraphqlHandlerOptions{
			Query:        schema.RootQuery,
			Mutation:     schema.RootMutation,
			RootObjectFn: rootObjectFn,
			Services:     services,
		},
	)

	http.Handle(gqlpath, graphqlHandler)
	fmt.Printf("Access the web app via browser at http://localhost:%d%s\n", port, gqlpath)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
