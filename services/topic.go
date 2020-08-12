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
	GetTopicsByPage(urlValues *url.Values) *[]models.Topic
	GetTopicById(id string) *models.TopicDetail
}

func NewTopicService(httpClient *utils.HttpClient, BaseURL string) *topicService {
	return &topicService{HttpClient: httpClient, BaseURL: BaseURL}
}
func (t *topicService) GetTopicsByPage(urlValues *url.Values) *[]models.Topic {
	base, err := url.Parse(t.BaseURL + "/topics")
	if err != nil {
		fmt.Println("Get topics by page error. reason: parse url error.", err)
		return &[]models.Topic{}
	}
	base.RawQuery = urlValues.Encode()
	body, err := t.HttpClient.Get(base.String())
	if err != nil {
		fmt.Println(err)
		return &[]models.Topic{}
	}
	return body.(*[]models.Topic)
}

func (t *topicService) GetTopicById(id string) *models.TopicDetail {
	endpoint := t.BaseURL + "/topic/" + id
	body, err := t.HttpClient.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return &models.TopicDetail{}
	}
	return body.(*models.TopicDetail)
}
