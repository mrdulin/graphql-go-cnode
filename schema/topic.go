package schema

import (
	"net/url"
	"strconv"

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

type RecentTopic struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	LastReplyAt string `json:"last_reply_at"`
	Author      User   `json:"author"`
}

var TopicTabEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "TopicTab",
	Description: "The category of topic",
	Values: graphql.EnumValueConfigMap{
		"ASK":   &graphql.EnumValueConfig{Value: "ask"},
		"SHARE": &graphql.EnumValueConfig{Value: "share"},
		"JOB":   &graphql.EnumValueConfig{Value: "job"},
		"GOOD":  &graphql.EnumValueConfig{Value: "good"},
	},
})

var TopicType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Topic",
	Description: "This is topic",
	Fields: graphql.Fields{
		"id":            &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"author_id":     &graphql.Field{Type: graphql.String},
		"tab":           &graphql.Field{Type: TopicTabEnum},
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
		var val string
		switch v := v.(type) {
		case int:
			val = strconv.Itoa(v)
		case string:
			val = v
		}
		urlValues.Add(k, val)
	}
	base.RawQuery = urlValues.Encode()
	body, err := utils.RequestGet(base.String())
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
	body, err := utils.RequestGet(url)
	if err != nil {
		return &TopicDetail{}, nil
	}
	return body.(utils.Response).Data, nil
}
