package services

import (
	"fmt"
	"net/url"

	"github.com/mrdulin/graphql-go-cnode/models"
	"github.com/mrdulin/graphql-go-cnode/utils"
)

type messageService struct {
	HttpClient utils.IHttpClient
	BaseURL    string
}

type MessageService interface {
	GetMessages(accesstoken, mdrender string) interface{}
	GetUnreadMessageCount(accesstoken string) interface{}
	MarkAll(accesstoken string) interface{}
}

func NewMessageService(httpClient utils.IHttpClient, BaseURL string) *messageService {
	return &messageService{HttpClient: httpClient, BaseURL: BaseURL}
}

func (m *messageService) GetMessages(accesstoken, mdrender string) interface{} {
	base, err := url.Parse(m.BaseURL + "/messages")
	res := models.GetMessagesResponse{}
	if err != nil {
		fmt.Println("Get messages error: parse url.", err)
		return &res
	}
	urlValues := url.Values{}
	urlValues.Add("accesstoken", accesstoken)
	urlValues.Add("mdrender", mdrender)
	base.RawQuery = urlValues.Encode()
	body, err := m.HttpClient.Get(base.String())
	if err != nil {
		fmt.Println(err)
		return &res
	}
	return body
}

func (m *messageService) GetUnreadMessageCount(accesstoken string) interface{} {
	endpoint := m.BaseURL + "/message/count?accesstoken=" + accesstoken
	body, err := m.HttpClient.Get(endpoint)
	var count int
	if err != nil {
		fmt.Println(err)
		return count
	}
	return body
}

func (m *messageService) MarkAll(accesstoken string) interface{} {
	endpoint := m.BaseURL + "/message/mark_all"
	body, err := m.HttpClient.Post(endpoint, &models.MarkAllRequest{AccessToken: accesstoken})
	res := models.MarkAllMessagesResponse{}
	if err != nil {
		fmt.Println(err)
		return &res
	}
	return body
}
