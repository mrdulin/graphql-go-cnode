package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/mrdulin/graphql-go-cnode/utils"
)

var ReplyBaseFields = graphql.Fields{
	"id":        &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
	"content":   &graphql.Field{Type: graphql.String},
	"ups":       &graphql.Field{Type: graphql.NewList(graphql.ID)},
	"create_at": &graphql.Field{Type: graphql.String},
}

var ReplyBaseType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ReplyBase",
	Description: "This is base information of a reply",
	Fields:      ReplyBaseFields,
})

var ReplyType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Reply",
	Description: "This is user's reply for a post",
	Fields: utils.MergeGraphqlFields(ReplyBaseFields, graphql.Fields{
		"reply_id": &graphql.Field{Type: graphql.ID},
		"is_uped":  &graphql.Field{Type: graphql.Boolean},
		"author": &graphql.Field{
			Type:    UserType,
			Resolve: AuthorResolver,
		},
	}),
})

func RepliesResolver(p graphql.ResolveParams) (interface{}, error) {
	source, _ := p.Source.(map[string]interface{})
	return source["replies"], nil
}
