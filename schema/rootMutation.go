package schema

import "github.com/graphql-go/graphql"

// RootMutation root mutation for every HTTP post request
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"validateAccessToken": &graphql.Field{
			Type:        AccessTokenValidationType,
			Description: "Validate accessToken",
			Args: graphql.FieldConfigArgument{
				"accessToken": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: AccessTokenValidationResolver,
		},

		"markAllMessages": &graphql.Field{
			Type:        MarkAllResponseType,
			Description: "Mark all messages",
			Args: graphql.FieldConfigArgument{
				"accessToken": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: MarkAllMessagesResolver,
		},
	},
})
