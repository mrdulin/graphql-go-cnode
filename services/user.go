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
	GetUserDetailByLoginname(name string) *models.UserDetail
	ValidateAccessToken(token string) *models.UserEntity
}

func NewUserService(httpClient utils.IHttpClient, BaseURL string) *userService {
	return &userService{HttpClient: httpClient, BaseURL: BaseURL}
}

func (u *userService) GetUserDetailByLoginname(name string) *models.UserDetail {
	url := u.BaseURL + "/user/" + name
	body, err := u.HttpClient.Get(url)
	userDetail := models.UserDetail{}
	if err != nil {
		fmt.Println(err)
		return &userDetail
	}
	u.HttpClient.Decode(body, &userDetail)
	return &userDetail
}

func (u *userService) ValidateAccessToken(token string) *models.UserEntity {
	url := u.BaseURL + "/accesstoken"
	body, err := u.HttpClient.Post(url, &models.ValidateAccessTokenRequest{AccessToken: token})
	userEntity := models.UserEntity{}
	if err != nil {
		fmt.Println(err)
		return &userEntity
	}
	u.HttpClient.Decode(body, &userEntity)
	return &userEntity
}
