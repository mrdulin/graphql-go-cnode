package services

import "github.com/mrdulin/graphql-go-cnode/models"

type userService struct{}

type UserService interface {
	GetUserDetailByLoginname(name string) *models.UserDetail
}

func NewUserService() *userService {
	return &userService{}
}

func (u *userService) GetUserDetailByLoginname(name string) *models.UserDetail {
	return &models.UserDetail{}
}
