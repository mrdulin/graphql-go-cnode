package schema

import (
	"fmt"
	"net/url"

	"github.com/graphql-go/graphql"
	utils "github.com/mrdulin/graphql-go-cnode/utils"
)

type TopicBase struct {
	ID          string `json:"id"`
	AuthorID    string `json:"author_id"`
	Tab         string `json:"tab"`
	Content     string `json:"content"`
	Title       string `json:"title"`
	LastReplyAt string `json:"last_reply_at"`
	Good        bool   `json:"good"`
	Top         bool   `json:"top"`
	ReplyCount  int    `json:"reply_count"`
	VisitCount  int    `json:"visit_count"`
	CreateAt    string `json:"create_at"`
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
	Name:        "Topic",
	Description: "This is topic",
	Fields: graphql.Fields{
		"id":            &graphql.Field{Type: graphql.String},
		"author_id":     &graphql.Field{Type: graphql.String},
		"tab":           &graphql.Field{Type: graphql.String},
		"content":       &graphql.Field{Type: graphql.String},
		"title":         &graphql.Field{Type: graphql.String},
		"last_reply_at": &graphql.Field{Type: graphql.String},
		"good":          &graphql.Field{Type: graphql.Boolean},
		"top":           &graphql.Field{Type: graphql.Int},
		"reply_count":   &graphql.Field{Type: graphql.Int},
		"visit_count":   &graphql.Field{Type: graphql.Int},
		"create_at":     &graphql.Field{Type: graphql.String},
		"author": &graphql.Field{
			Type:    UserType,
			Resolve: AuthorResolver,
		},
		"replies": &graphql.Field{
			Type:    graphql.NewList(ReplyType),
			Resolve: RepliesResolver,
		},
	},
})

func TopicsResolver(params graphql.ResolveParams) (interface{}, error) {
	base, err := url.Parse("https://cnodejs.org/api/v1/topics")
	if err != nil {
		return nil, err
	}
	urlValues := url.Values{}
	for k, v := range params.Args {
		// TODO: validation value
		fmt.Println("k:", k, "v:", v)
		urlValues.Add(k, v.(string))
	}
	base.RawQuery = urlValues.Encode()
	body, err := utils.Request(base.String())
	if err != nil {
		return Topic{}, nil
	}
	return body.(utils.Response).Data, nil
}

func TopicResolver(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(string)
	if !ok {
		return &TopicDetail{}, nil
	}
	url := "https://cnodejs.org/api/v1/topic/" + id
	body, err := utils.Request(url)
	if err != nil {
		return &TopicDetail{}, nil
	}
	return body.(utils.Response).Data, nil
}
