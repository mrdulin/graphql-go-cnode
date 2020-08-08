package schema

import (
	"github.com/graphql-go/graphql"
)

type User struct {
	Loginname string `json:"loginname"`
	AvatarURL string `json:"avatar_url"`
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "This respresents an Author",
	Fields: graphql.Fields{
		"loginname":  &graphql.Field{Type: graphql.String},
		"avatar_url": &graphql.Field{Type: graphql.String},
	},
})

func AuthorResolver(p graphql.ResolveParams) (interface{}, error) {
	source, _ := p.Source.(map[string]interface{})
	author := source["author"]
	return author, nil
}
