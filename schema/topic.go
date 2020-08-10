package schema

import (
	"net/url"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/mrdulin/graphql-go-cnode/models"
	utils "github.com/mrdulin/graphql-go-cnode/utils"
)

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
		"id":            &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
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

var RecentTopicType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RecentTopic",
	Description: "Recent topic of an user",
	Fields: graphql.Fields{
		"id":            &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"title":         &graphql.Field{Type: graphql.String},
		"last_reply_at": &graphql.Field{Type: graphql.String},
		"author":        &graphql.Field{Type: UserType},
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
		return &models.Topic{}, nil
	}
	return body.(utils.Response).Data, nil
}

func TopicResolver(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(string)
	if !ok {
		return &models.TopicDetail{}, nil
	}
	url := "https://cnodejs.org/api/v1/topic/" + id
	body, err := utils.RequestGet(url)
	if err != nil {
		return &models.TopicDetail{}, nil
	}
	return body.(utils.Response).Data, nil
}
