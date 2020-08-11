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
	if err != nil {
		fmt.Println("Get user detail by login name error.", err)
		return &models.UserDetail{}
	}
	return body.(utils.Response).Data
}

func (u *userService) ValidateAccessToken(token string) interface{} {
	url := u.BaseURL + "/accesstoken"
	body, err := u.HttpClient.Post(url, map[string]interface{}{"accesstoken": token})
	if err != nil {
		fmt.Println("Validate access token error.", err)
		return &models.AccessTokenValidation{}
	}
	return body
}
