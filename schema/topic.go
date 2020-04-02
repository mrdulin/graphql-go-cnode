package schema

import (
	"github.com/graphql-go/graphql"
	"time"
)

type Topic struct {
	ID          string    `json:"id"`
	AuthorID    string    `json:"author_id"`
	Tab         string    `json:"tab"`
	Content     string    `json:"content"`
	Title       string    `json:"title"`
	LastReplyAt time.Time `json:"last_reply_at"`
	Good        bool      `json:"good"`
	Top         bool      `json:"top"`
	ReplyCount  int       `json:"reply_count"`
	VisitCount  int       `json:"visit_count"`
	CreateAt    time.Time `json:"create_at"`
	Author      struct {
		Loginname string `json:"loginname"`
		AvatarURL string `json:"avatar_url"`
	} `json:"author"`
}

var TopicType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Topic",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.String},
		"author_id": &graphql.Field{Type: graphql.String},
		"tab":       &graphql.Field{Type: graphql.String},
		"content":   &graphql.Field{Type: graphql.String},
		"title":     &graphql.Field{Type: graphql.String},
	},
})
