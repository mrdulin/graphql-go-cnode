package models

import "github.com/mrdulin/graphql-go-cnode/utils"

type User struct {
	Loginname string `json:"loginname"`
	AvatarURL string `json:"avatar_url"`
}

type UserDetail struct {
	User
	GithubUsername string        `json:"githubUsername"`
	CreateAt       string        `json:"create_at"`
	Score          int           `json:"score"`
	RecentTopics   []RecentTopic `json:"recent_topics"`
}

type ValidateAccessTokenResponse struct {
	utils.ResponseStatus
	User
	ID string `json:"id"`
}
