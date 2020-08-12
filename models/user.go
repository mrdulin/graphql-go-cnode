package models

type User struct {
	Loginname string `json:"loginname"`
	AvatarURL string `json:"avatar_url" mapstructure:"avatar_url"`
}

type UserDetail struct {
	User           `mapstructure:",squash"`
	GithubUsername string        `json:"githubUsername"`
	CreateAt       string        `json:"create_at" mapstructure:"create_at"`
	Score          int           `json:"score"`
	RecentTopics   []RecentTopic `json:"recent_topics" mapstructure:"recent_topics"`
}

type UserEntity struct {
	User `mapstructure:",squash"`
	ID   string `json:"id"`
}

type ValidateAccessTokenRequest struct {
	AccessToken string `json:"accesstoken"`
}
