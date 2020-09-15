package main_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"github.com/mrdulin/graphql-go-cnode/controllers"
	"github.com/mrdulin/graphql-go-cnode/models"
	"github.com/mrdulin/graphql-go-cnode/schema"
	svcs "github.com/mrdulin/graphql-go-cnode/services"
	"github.com/mrdulin/graphql-go-cnode/utils"
	assert2 "github.com/stretchr/testify/assert"
)

var (
	services       *svcs.Container
	envFetcher     utils.EnvFetcher
	gqlpath        = "/graphql"
	port           = 3000
	graphqlHandler *handler.Handler
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
	rootObjectFn := controllers.NewRootObjectFn(services)
	graphqlHandler = controllers.NewGraphqlHandler(
		controllers.GraphqlHandlerOptions{
			Query:        schema.RootQuery,
			Mutation:     schema.RootMutation,
			RootObjectFn: rootObjectFn,
			Services:     services,
		},
	)
}

// API testings
func TestGetUserDetailByLoginname(t *testing.T) {
	w := httptest.NewRecorder()
	assert := assert2.New(t)
	t.Run("should get user detail by login name success", func(t *testing.T) {
		query := `
			query {
			  user(loginname: "mrdulin") {
				avatar_url
				create_at
				githubUsername
				loginname
				score
			  }
			}
		`
		reader := strings.NewReader(query)
		r, _ := http.NewRequest(http.MethodPost, gqlpath, reader)
		r.Header.Set("Content-Type", "application/graphql")
		graphqlHandler.ServeHTTP(w, r)
		assert.Equal(200, w.Code, "should return status code 200")
		var resp struct {
			Data struct {
				User models.UserDetail
			}
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatal(err)
		}
		//t.Logf("body: %+v", w.Body.String())
		assert.NotNil(resp.Data.User.AvatarURL)
		assert.Equal("mrdulin", resp.Data.User.Loginname)
		assert.JSONEq(`{
        	"data": {
        		"user": {
        			"avatar_url": "https://avatars2.githubusercontent.com/u/17866683?v=4\u0026s=120",
        			"create_at": "2017-04-02T14:03:57.934Z",
        			"githubUsername": "mrdulin",
        			"loginname": "mrdulin",
        			"score": 50
        		}
        	}
        }`, w.Body.String(), "should get correct data")
	})
}
