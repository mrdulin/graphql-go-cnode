package schema

import (
	"github.com/graphql-go/graphql"
)

// RootQuery root query for every HTTP GET request
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"topics": &graphql.Field{
			Type: graphql.NewList(TopicType),
			Args: graphql.FieldConfigArgument{
				"page":     &graphql.ArgumentConfig{Type: graphql.Int},
				"tab":      &graphql.ArgumentConfig{Type: TopicTabEnum},
				"limit":    &graphql.ArgumentConfig{Type: graphql.Int},
				"mdrender": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: TopicsResolver,
		},

		"topic": &graphql.Field{
			Type: TopicDetailType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: TopicResolver,
		},

		"user": &graphql.Field{
			Type:        UserDetailType,
			Description: "Get user detail by login name",
			Args: graphql.FieldConfigArgument{
				"loginname": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: UserDetailResolver,
		},

		"messages": &graphql.Field{
			Type:        MessagesType,
			Description: "Contain unread and read messages",
			Args: graphql.FieldConfigArgument{
				"accessToken": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"mdrender":    &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: "true"},
			},
			Resolve: MessagesResolver,
		},
	},
})
