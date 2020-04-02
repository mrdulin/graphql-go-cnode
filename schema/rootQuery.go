package schema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/graphql-go/graphql"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func request(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data Response
	json.Unmarshal(body, &data)
	return data, nil
}

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"topics": &graphql.Field{
			Type: graphql.NewList(TopicType),
			Args: graphql.FieldConfigArgument{
				"limit": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				base, err := url.Parse("https://cnodejs.org/api/v1/topics")
				if err != nil {
					return nil, err
				}
				urlValues := url.Values{}
				for k, v := range params.Args {
					// TODO: validation value
					fmt.Println("k:", k, "v:", v)
					urlValues.Add(k, v.(string))
				}
				base.RawQuery = urlValues.Encode()
				fmt.Printf("Encoded URL is %q\n", base.String())
				body, err := request(base.String())
				fmt.Println(body.(Response).Success)
				if err != nil {
					return Topic{}, nil
				}
				return body.(Response).Data, nil
			},
		},
	},
})
