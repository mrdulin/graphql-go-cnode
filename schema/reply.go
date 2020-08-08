package schema

import (
	"time"

	"github.com/graphql-go/graphql"
)

type Reply struct {
	ID       string      `json:"id"`
	Author   User        `json:"author"`
	Content  string      `json:"content"`
	Ups      []string    `json:"ups"`
	CreateAt time.Time   `json:"create_at"`
	ReplyID  interface{} `json:"reply_id"`
	IsUped   bool        `json:"is_uped"`
}

var ReplyType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Reply",
	Description: "This is user's reply for a post",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"author": &graphql.Field{
			Type:    UserType,
			Resolve: AuthorResolver,
		},
	},
})

func RepliesResolver(p graphql.ResolveParams) (interface{}, error) {
	source, _ := p.Source.(map[string]interface{})
	return source["replies"], nil
}
