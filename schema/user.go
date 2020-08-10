package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/mrdulin/graphql-go-cnode/models"
	"github.com/mrdulin/graphql-go-cnode/services"
	"github.com/mrdulin/graphql-go-cnode/utils"
)

var UserBaseFields = graphql.Fields{
	"loginname":  &graphql.Field{Type: graphql.String},
	"avatar_url": &graphql.Field{Type: graphql.String},
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "This respresents an user",
	Fields:      UserBaseFields,
})

var UserDetailType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserDetail",
	Description: "This respresents an user detail",
	Fields: utils.MergeGraphqlFields(UserBaseFields, graphql.Fields{
		"githubUsername": &graphql.Field{Type: graphql.String},
		"create_at":      &graphql.Field{Type: graphql.String},
		"score":          &graphql.Field{Type: graphql.Int},
		"recent_topics":  &graphql.Field{Type: graphql.NewList(RecentTopicType)},
	}),
})

var AccessTokenValidationType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "AccessTokenValidation",
	Description: "The response type for validating accessToken",
	Fields: utils.MergeGraphqlFields(UserBaseFields, graphql.Fields{
		"id":      &graphql.Field{Type: graphql.ID},
		"success": &graphql.Field{Type: graphql.Boolean},
	}),
})

func UserDetailResolver(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	container := rootValue["services"].(*services.Container)
	loginname, ok := params.Args["loginname"].(string)
	if !ok {
		return &models.UserDetail{}, nil
	}
	return container.UserService.GetUserDetailByLoginname(loginname), nil
}

func AuthorResolver(p graphql.ResolveParams) (interface{}, error) {
	source, _ := p.Source.(map[string]interface{})
	author := source["author"]
	return author, nil
}

func AccessTokenValidationResolver(p graphql.ResolveParams) (interface{}, error) {
	accessToken, ok := p.Args["accessToken"].(string)
	if !ok {
		return &models.AccessTokenValidation{}, nil
	}
	rootValue := p.Info.RootValue.(map[string]interface{})
	container := rootValue["services"].(*services.Container)
	return container.UserService.ValidateAccessToken(accessToken), nil
}
