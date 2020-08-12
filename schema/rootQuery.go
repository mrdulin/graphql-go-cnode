package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/mrdulin/graphql-go-cnode/utils"
)

// RootQuery root query for every HTTP GET request
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:   "RootQuery",
	Fields: utils.MergeGraphqlFields(MessageQueryFields, TopicQueryFields, UserQueryFields),
})
