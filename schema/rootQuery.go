package schema

import (
	"github.com/graphql-go/graphql"
)

// RootQuery graphql root query
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
			Type: TopicType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: TopicResolver,
		},

		"user": &graphql.Field{
			Type:        UserType,
			Description: "get user detail by login name",
			Args: graphql.FieldConfigArgument{
				"loginname": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: UserDetailResolver,
		},

		"validateAccessToken": &graphql.Field{
			Type:        AccessTokenValidationType,
			Description: "validate accessToken",
			Args: graphql.FieldConfigArgument{
				"accessToken": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: AccessTokenValidationResolver,
		},
	},
})
