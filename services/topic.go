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
	if err != nil {
		fmt.Println("Get topics by page error. reason: parse url error.", err)
		return &[]models.Topic{}
	}
	base.RawQuery = urlValues.Encode()
	body, err := t.HttpClient.Get(base.String())
	if err != nil {
		fmt.Println("Get topics by page error. reason: http get error.", err)
		return &[]models.Topic{}
	}
	return body.(utils.Response).Data
}

func (t *topicService) GetTopicById(id string) interface{} {
	endpoint := t.BaseURL + "/topic/" + id
	body, err := t.HttpClient.Get(endpoint)
	if err != nil {
		fmt.Println("Get topic by Id error.", err)
		return &models.TopicDetail{}
	}
	return body.(utils.Response).Data
}
