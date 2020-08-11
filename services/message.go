package services

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/mrdulin/graphql-go-cnode/models"
	"github.com/mrdulin/graphql-go-cnode/utils"
)

type MarkAllRequest struct {
	AccessToken string `json:"accesstoken"`
}

type messageService struct {
	HttpClient utils.IHttpClient
	BaseURL    string
}

type MessageService interface {
	GetMessages(accesstoken, mdrender string) interface{}
	GetUnreadMessageCount(accesstoken string) interface{}
	MarkAll(accesstoken string) *models.MarkAllMessagesResponse
}

func NewMessageService(httpClient utils.IHttpClient, BaseURL string) *messageService {
	return &messageService{HttpClient: httpClient, BaseURL: BaseURL}
}

func (m *messageService) GetMessages(accesstoken, mdrender string) interface{} {
	base, err := url.Parse(m.BaseURL + "/messages")
	if err != nil {
		fmt.Println("Get messages error. reason: parse url error.", err)
		return &models.Messages{}
	}
	urlValues := url.Values{}
	urlValues.Add("accesstoken", accesstoken)
	urlValues.Add("mdrender", mdrender)
	base.RawQuery = urlValues.Encode()
	body, err := m.HttpClient.Get(base.String())
	if err != nil {
		fmt.Println("Get messages error. reason: HTTP request error.", err)
		return &models.Messages{}
	}
	res := body.(utils.Response)
	if !res.Success {
		fmt.Println("Get messages error. reason: API error")
		return &models.Messages{}
	}
	return res.Data
}

func (m *messageService) GetUnreadMessageCount(accesstoken string) interface{} {
	endpoint := m.BaseURL + "/messages?accesstoken=" + accesstoken
	body, err := m.HttpClient.Get(endpoint)
	if err != nil {
		fmt.Println("Get unread message")
		return 0
	}
	res := body.(utils.Response)
	if !res.Success {
		fmt.Println("Get unread message count. reason: API error.")
		return 0
	}
	return res.Data
}

func (m *messageService) MarkAll(accesstoken string) *models.MarkAllMessagesResponse {
	endpoint := m.BaseURL + "/message/mark_all"
	body, err := m.HttpClient.Post(endpoint, &MarkAllRequest{AccessToken: accesstoken})
	res := models.MarkAllMessagesResponse{}
	if err != nil {
		fmt.Println("Mark all messages error.", err)
		return &res
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("json.Unmarshal error.")
		return &res
	}
	if !res.Success {
		fmt.Println("Mark all messages count. reason: API error.")
		return &res
	}
	return &res
}
