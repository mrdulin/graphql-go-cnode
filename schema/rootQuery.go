package schema

import (
	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"topics": &graphql.Field{
			Type: graphql.NewList(TopicType),
			Args: graphql.FieldConfigArgument{
				"limit": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: TopicsResolver,
		},

		// "topic": &graphql.Field{
		// 	Type
		// },
	},
})
