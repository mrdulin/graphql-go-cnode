package schema

import (
	"fmt"

	"github.com/mrdulin/graphql-go-cnode/utils"

	"github.com/graphql-go/graphql"
	"github.com/mrdulin/graphql-go-cnode/models"
	"github.com/mrdulin/graphql-go-cnode/services"
)

var MessageIdField = graphql.Fields{
	"id": &graphql.Field{
		Type: graphql.NewNonNull(graphql.ID),
	},
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

var MessageQueryFields = graphql.Fields{
	"messages": &graphql.Field{
		Type:        MessagesType,
		Description: "Contain unread and read messages",
		Args: graphql.FieldConfigArgument{
			"accessToken": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"mdrender":    &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: "true"},
		},
		Resolve: MessagesResolver,
	},

	"MessageCount": &graphql.Field{
		Type: graphql.Int,
		Name: "Get total count of messages of an user",
		Args: graphql.FieldConfigArgument{
			"accessToken": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: MessageCountResolver,
	},
}

var MessageMutationFields = graphql.Fields{
	"markAllMessages": &graphql.Field{
		Type:        MarkAllResponseType,
		Description: "Mark all messages",
		Args: graphql.FieldConfigArgument{
			"accessToken": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: MarkAllMessagesResolver,
	},
}

func MessagesResolver(p graphql.ResolveParams) (interface{}, error) {
	rootValue := p.Info.RootValue.(map[string]interface{})
	container := rootValue["services"].(*services.Container)
	var ok bool
	accessToken, ok := p.Args["accessToken"].(string)
	if !ok {
		fmt.Println("resolver params 'accessToken' type cast error.")
		return &models.GetMessagesResponse{}, nil
	}
	mdrender, ok := p.Args["mdrender"].(string)
	if !ok {
		fmt.Println("resolver params 'mdrender' type cast error.")
		return &models.GetMessagesResponse{}, nil
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

func MessageCountResolver(p graphql.ResolveParams) (interface{}, error) {
	return 0, nil
}

// If return type is interface{} for the method of MessageService, this resolver is unnecessary
//func HasReadMessagesResolver(p graphql.ResolveParams) (interface{}, error) {
//	source, ok := p.Source.(map[string]interface{})
//	if !ok {
//		fmt.Println("type cast p.Source.(*models.GetMessagesResponse) error in HasReadMessagesResolver")
//		return models.GetMessagesResponse{}.HasReadMessages, nil
//	}
//	return source["has_read_messages"], nil
//}
