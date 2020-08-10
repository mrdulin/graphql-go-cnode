package services

import (
	"fmt"
	"net/url"

	"github.com/mrdulin/graphql-go-cnode/models"
	"github.com/mrdulin/graphql-go-cnode/utils"
)

type topicService struct {
	RequestGet utils.RequestGetter
	BaseURL    string
}

type TopicService interface {
	GetTopicsByPage(urlValues *url.Values) interface{}
	GetTopicById(id string) interface{}
}

func NewTopicService(requestGet utils.RequestGetter, BaseURL string) *topicService {
	return &topicService{RequestGet: requestGet, BaseURL: BaseURL}
}
func (t *topicService) GetTopicsByPage(urlValues *url.Values) interface{} {
	base, err := url.Parse(t.BaseURL + "/topics")
	if err != nil {
		fmt.Println("Get topics by page error. reason: parse url error.", err)
		return &[]models.Topic{}
	}
	base.RawQuery = urlValues.Encode()
	body, err := t.RequestGet(base.String())
	if err != nil {
		fmt.Println("Get topics by page error. reason: http get error.", err)
		return &[]models.Topic{}
	}
	return body.(utils.Response).Data
}

func (t *topicService) GetTopicById(id string) interface{} {
	endpoint := t.BaseURL + "/topic/" + id
	body, err := t.RequestGet(endpoint)
	if err != nil {
		fmt.Println("Get topic by Id error.", err)
		return &models.TopicDetail{}
	}
	return body.(utils.Response).Data
}
