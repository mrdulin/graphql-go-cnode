package schema

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
)

type TopicBase struct {
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
}

type Topic struct {
	TopicBase
	Author User `json:"author"`
}

type TopicDetail struct {
	Topic
	Replies []Reply `json:"replies"`
}

var TopicType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Topic",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.String},
		"author_id": &graphql.Field{Type: graphql.String},
		"tab":       &graphql.Field{Type: graphql.String},
		"content":   &graphql.Field{Type: graphql.String},
		"title":     &graphql.Field{Type: graphql.String},
		"author": &graphql.Field{
			Type:    UserType,
			Resolve: AuthorResolver,
		},
	},
})

func PrintPretty(x interface{}) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}

func AuthorResolver(p graphql.ResolveParams) (interface{}, error) {
	source, _ := p.Source.(map[string]interface{})
	author, _ := source["author"]
	return author, nil
}
