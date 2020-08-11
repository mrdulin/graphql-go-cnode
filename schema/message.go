package schema

import (
	"fmt"

	"github.com/mrdulin/graphql-go-cnode/utils"

	"github.com/graphql-go/graphql"
	"github.com/mrdulin/graphql-go-cnode/models"
	"github.com/mrdulin/graphql-go-cnode/services"
)

var MessageIdField = graphql.Fields{
	"id": &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
}

var MessageIdType = graphql.NewObject(graphql.ObjectConfig{
	Name:   "MessageId",
	Fields: MessageIdField,
})

var MessageType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Message",
	Description: "This is message",
	Fields: utils.MergeGraphqlFields(MessageIdField, graphql.Fields{
		"type":      &graphql.Field{Type: graphql.String},
		"has_read":  &graphql.Field{Type: graphql.Boolean},
		"author":    &graphql.Field{Type: UserType},
		"topic":     &graphql.Field{Type: TopicBaseType},
		"reply":     &graphql.Field{Type: ReplyBaseType},
		"create_at": &graphql.Field{Type: graphql.String},
	}),
})

var MessagesType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Messages",
	Description: "These are read and unread messages",
	Fields: graphql.Fields{
		"has_read_messages":    &graphql.Field{Type: graphql.NewList(MessageType)},
		"hasnot_read_messages": &graphql.Field{Type: graphql.NewList(MessageType)},
	},
})

var MarkAllResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "MarkedMessagesRepsonse",
	Description: "marked messages API response",
	Fields: graphql.Fields{
		"marked_msgs": &graphql.Field{Type: graphql.NewList(MessageIdType)},
	},
})

func MessagesResolver(p graphql.ResolveParams) (interface{}, error) {
	rootValue := p.Info.RootValue.(map[string]interface{})
	container := rootValue["services"].(*services.Container)
	var ok bool
	accessToken, ok := p.Args["accessToken"].(string)
	if !ok {
		fmt.Println("resolver params 'accessToken' type cast error.")
		return &models.Messages{}, nil
	}
	mdrender, ok := p.Args["mdrender"].(string)
	if !ok {
		fmt.Println("resolver params 'mdrender' type cast error.")
		return &models.Messages{}, nil
	}
	return container.MessageService.GetMessages(accessToken, mdrender), nil
}

func MarkAllMessagesResolver(p graphql.ResolveParams) (interface{}, error) {
	rootValue := p.Info.RootValue.(map[string]interface{})
	container := rootValue["services"].(*services.Container)
	accessToken, ok := p.Args["accessToken"].(string)
	if !ok {
		fmt.Println("resolver params 'accessToken' type cast error.")
		return &models.MarkAllMessagesResponse{}, nil
	}
	return container.MessageService.MarkAll(accessToken), nil
}
