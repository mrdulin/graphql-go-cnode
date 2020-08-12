package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/mrdulin/graphql-go-cnode/utils"
)

// RootMutation root mutation for every HTTP post request
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:   "RootMutation",
	Fields: utils.MergeGraphqlFields(UserMutationFields, MessageMutationFields),
})
