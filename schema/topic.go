package schema

import (
	"fmt"
	"net/url"
	"time"

	"github.com/graphql-go/graphql"
	utils "github.com/mrdulin/graphql-go-cnode/utils"
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
	fmt.Printf("Encoded URL is %q\n", base.String())
	body, err := utils.Request(base.String())
	fmt.Println(body.(utils.Response).Success)
	if err != nil {
		return Topic{}, nil
	}
	return body.(utils.Response).Data, nil
}

func AuthorResolver(p graphql.ResolveParams) (interface{}, error) {
	source, _ := p.Source.(map[string]interface{})
	author := source["author"]
	return author, nil
}
