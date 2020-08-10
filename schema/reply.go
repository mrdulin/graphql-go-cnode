package schema

import (
	"github.com/graphql-go/graphql"
)

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
