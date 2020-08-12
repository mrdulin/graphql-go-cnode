package services

import (
	"fmt"

	"github.com/mrdulin/graphql-go-cnode/models"
	"github.com/mrdulin/graphql-go-cnode/utils"
)

type userService struct {
	HttpClient utils.IHttpClient
	BaseURL    string
}

type UserService interface {
	GetUserDetailByLoginname(name string) interface{}
	ValidateAccessToken(token string) interface{}
}

func NewUserService(httpClient utils.IHttpClient, BaseURL string) *userService {
	return &userService{HttpClient: httpClient, BaseURL: BaseURL}
}

func (u *userService) GetUserDetailByLoginname(name string) interface{} {
	url := u.BaseURL + "/user/" + name
	body, err := u.HttpClient.Get(url)
	userDetail := models.UserDetail{}
	if err != nil {
		fmt.Println(err)
		return &userDetail
	}
	return body
}

func (u *userService) ValidateAccessToken(token string) interface{} {
	url := u.BaseURL + "/accesstoken"
	body, err := u.HttpClient.Post(url, &models.ValidateAccessTokenRequest{AccessToken: token})
	userEntity := models.UserEntity{}
	if err != nil {
		fmt.Println(err)
		return &userEntity
	}
	return body
}
