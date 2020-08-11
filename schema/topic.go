package schema

import (
	"net/url"
	"strconv"

	"github.com/mrdulin/graphql-go-cnode/utils"

	"github.com/graphql-go/graphql"
	"github.com/mrdulin/graphql-go-cnode/models"
	"github.com/mrdulin/graphql-go-cnode/services"
)

var TopicTabEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "TopicTab",
	Description: "The category of topic",
	Values: graphql.EnumValueConfigMap{
		"ASK":   &graphql.EnumValueConfig{Value: models.TOPIC_TAB_ASK},
		"SHARE": &graphql.EnumValueConfig{Value: models.TOPIC_TAB_SHARE},
		"JOB":   &graphql.EnumValueConfig{Value: models.TOPIC_TAB_JOB},
		"GOOD":  &graphql.EnumValueConfig{Value: models.TOPIC_TAB_GOOD},
	},
})

var TopicBaseFields = graphql.Fields{
	"id":            &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
	"title":         &graphql.Field{Type: graphql.String},
	"last_reply_at": &graphql.Field{Type: graphql.String},
}

var TopicFields = utils.MergeGraphqlFields(TopicBaseFields, graphql.Fields{
	"author_id":   &graphql.Field{Type: graphql.String},
	"tab":         &graphql.Field{Type: TopicTabEnum},
	"content":     &graphql.Field{Type: graphql.String},
	"good":        &graphql.Field{Type: graphql.Boolean},
	"top":         &graphql.Field{Type: graphql.Int},
	"reply_count": &graphql.Field{Type: graphql.Int},
	"visit_count": &graphql.Field{Type: graphql.Int},
	"create_at":   &graphql.Field{Type: graphql.String},
	"author": &graphql.Field{
		Type:    UserType,
		Resolve: AuthorResolver,
	},
})

var TopicBaseType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "TopicBase",
	Description: "This is base information of a topic",
	Fields:      TopicBaseFields,
})

var TopicType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Topic",
	Description: "This is topic",
	Fields:      TopicFields,
})

var TopicDetailType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "TopicDetail",
	Description: "This is topic detail",
	Fields: utils.MergeGraphqlFields(TopicFields, graphql.Fields{
		"author": &graphql.Field{
			Type:    UserType,
			Resolve: AuthorResolver,
		},
		"replies": &graphql.Field{
			Type:    graphql.NewList(ReplyType),
			Resolve: RepliesResolver,
		},
	}),
})

var RecentTopicType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RecentTopic",
	Description: "Recent topic of an user",
	Fields:      TopicBaseFields,
})

func TopicsResolver(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	container := rootValue["services"].(*services.Container)

	urlValues := url.Values{}
	for k, v := range params.Args {
		// TODO: validate params
		var val string
		switch v := v.(type) {
		case int:
			val = strconv.Itoa(v)
		case string:
			val = v
		}
		urlValues.Add(k, val)
	}

	return container.TopicService.GetTopicsByPage(&urlValues), nil
}

func TopicResolver(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	container := rootValue["services"].(*services.Container)
	id, ok := params.Args["id"].(string)
	if !ok {
		return &models.TopicDetail{}, nil
	}
	return container.TopicService.GetTopicById(id), nil
}
