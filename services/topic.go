package services

import (
	"fmt"
	"net/url"

	"github.com/mrdulin/graphql-go-cnode/models"
	"github.com/mrdulin/graphql-go-cnode/utils"
)

type topicService struct {
	HttpClient *utils.HttpClient
	BaseURL    string
}

type TopicService interface {
	GetTopicsByPage(urlValues *url.Values) interface{}
	GetTopicById(id string) interface{}
}

func NewTopicService(httpClient *utils.HttpClient, BaseURL string) *topicService {
	return &topicService{HttpClient: httpClient, BaseURL: BaseURL}
}
func (t *topicService) GetTopicsByPage(urlValues *url.Values) interface{} {
	base, err := url.Parse(t.BaseURL + "/topics")
	res := []models.Topic{}
	if err != nil {
		fmt.Println("Get topics by page error: parse url.", err)
		return &res
	}
	base.RawQuery = urlValues.Encode()
	body, err := t.HttpClient.Get(base.String())
	if err != nil {
		fmt.Println(err)
		return &res
	}
	return body
}

func (t *topicService) GetTopicById(id string) interface{} {
	endpoint := t.BaseURL + "/topic/" + id
	body, err := t.HttpClient.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return body
	}
	return body
}
