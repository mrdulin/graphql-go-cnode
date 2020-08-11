package utils_test

import (
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/mrdulin/graphql-go-cnode/utils"
)

func TestMergeGraphqlFields(t *testing.T) {
	t.Run("should merge graphql fields correctly", func(t *testing.T) {
		a := graphql.Fields{"name": &graphql.Field{Type: graphql.String}}
		b := graphql.Fields{"email": &graphql.Field{Type: graphql.String}}
		got := utils.MergeGraphqlFields(a, b)
		want := graphql.Fields{
			"name":  &graphql.Field{Type: graphql.String},
			"email": &graphql.Field{Type: graphql.String},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("merge graphql fields error, got: %+v, want: %+v", got, want)
		}
		if !reflect.DeepEqual(a, graphql.Fields{"name": &graphql.Field{Type: graphql.String}}) {
			t.Errorf("should not mutate original map a")
		}
		if !reflect.DeepEqual(b, graphql.Fields{"email": &graphql.Field{Type: graphql.String}}) {
			t.Errorf("should not mutate original map b")
		}
	})
}

func TestMergeGraphqlFieldsConcurrency(t *testing.T) {
	t.Run("should merge graphql fields correctly", func(t *testing.T) {
		a := graphql.Fields{"name": &graphql.Field{Type: graphql.String}}
		b := graphql.Fields{"email": &graphql.Field{Type: graphql.String}}
		got := utils.MergeGraphqlFieldsConcurrency(a, b)
		want := graphql.Fields{
			"name":  &graphql.Field{Type: graphql.String},
			"email": &graphql.Field{Type: graphql.String},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("merge graphql fields error, got: %+v, want: %+v", got, want)
		}
		if !reflect.DeepEqual(a, graphql.Fields{"name": &graphql.Field{Type: graphql.String}}) {
			t.Errorf("should not mutate original map a")
		}
		if !reflect.DeepEqual(b, graphql.Fields{"email": &graphql.Field{Type: graphql.String}}) {
			t.Errorf("should not mutate original map b")
		}
	})
}

func BenchmarkMergeGraphqlFields(b *testing.B) {
	x := graphql.Fields{"name": &graphql.Field{Type: graphql.String}}
	y := graphql.Fields{"email": &graphql.Field{Type: graphql.String}}
	for i := 0; i < b.N; i++ {
		utils.MergeGraphqlFields(x, y)
	}
}

func BenchmarkMergeGraphqlFieldsConcurrency(b *testing.B) {
	x := graphql.Fields{"name": &graphql.Field{Type: graphql.String}}
	y := graphql.Fields{"email": &graphql.Field{Type: graphql.String}}
	for i := 0; i < b.N; i++ {
		utils.MergeGraphqlFieldsConcurrency(x, y)
	}
}
