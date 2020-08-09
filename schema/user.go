package schema

import (
	"github.com/graphql-go/graphql"
	utils "github.com/mrdulin/graphql-go-cnode/utils"
	"github.com/pkg/errors"
)

type User struct {
	Loginname string `json:"loginname"`
	AvatarURL string `json:"avatar_url"`
}

type UserDetail struct {
	User
	GithubUsername string        `json:"githubUsername"`
	CreateAt       string        `json:"create_at"`
	Score          int           `json:"score"`
	RecentTopics   []RecentTopic `json:"recent_topics"`
}

type AccessTokenValidation struct {
	utils.ResponseStatus
	User
	ID string `json:"id"`
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "This respresents an user",
	Fields: graphql.Fields{
		"loginname":  &graphql.Field{Type: graphql.String},
		"avatar_url": &graphql.Field{Type: graphql.String},
	},
})

var UserDetailType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserDetail",
	Description: "This respresents an user detail",
	Fields: graphql.Fields{
		// TODO: reuse fields definition of UserType here
		"loginname":  &graphql.Field{Type: graphql.String},
		"avatar_url": &graphql.Field{Type: graphql.String},

		"githubUsername": &graphql.Field{Type: graphql.String},
		"create_at":      &graphql.Field{Type: graphql.String},
		"score":          &graphql.Field{Type: graphql.Int},
		// TODO: recent_topics
	},
})

var AccessTokenValidationType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "AccessTokenValidation",
	Description: "The response type for validating accessToken",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.ID},
		"loginname":  &graphql.Field{Type: graphql.String},
		"avatar_url": &graphql.Field{Type: graphql.String},
		"success":    &graphql.Field{Type: graphql.Boolean},
	},
})

func UserDetailResolver(params graphql.ResolveParams) (interface{}, error) {
	loginname, ok := params.Args["loginname"].(string)
	if !ok {
		return &UserDetail{}, nil
	}
	url := "https://cnodejs.org/api/v1/user/" + loginname
	body, err := utils.RequestGet(url)
	if err != nil {
		return nil, err
	}
	return body.(utils.Response).Data, nil
}

func AuthorResolver(p graphql.ResolveParams) (interface{}, error) {
	source, _ := p.Source.(map[string]interface{})
	author := source["author"]
	return author, nil
}

func AccessTokenValidationResolver(p graphql.ResolveParams) (interface{}, error) {
	accessToken, ok := p.Args["accessToken"].(string)
	if !ok {
		return &AccessTokenValidation{}, nil
	}
	url := "https://cnodejs.org/api/v1/accesstoken"
	body, err := utils.RequestPost(url, map[string]interface{}{"accesstoken": accessToken})
	if err != nil {
		return nil, errors.Wrap(err, "utils.RequestPost")
	}
	return body, nil
}
