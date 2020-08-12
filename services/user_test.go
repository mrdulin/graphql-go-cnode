package services_test

import (
	"reflect"
	"testing"

	"github.com/mrdulin/graphql-go-cnode/mocks"
	"github.com/mrdulin/graphql-go-cnode/services"
)

func TestUserService_GetUserDetailByLoginname(t *testing.T) {
	t.Run("should get user detail by login name success", func(t *testing.T) {

		testHttpClient := new(mocks.MockedHttpClient)
		apiurl := "http://localhost/api"
		data := map[string]interface{}{
			"loginname": "mrdulin",
		}
		testHttpClient.On("Get", apiurl+"/user/mrdulin").Return(data, nil)
		userService := services.NewUserService(testHttpClient, apiurl)
		got := userService.GetUserDetailByLoginname("mrdulin")
		want := map[string]interface{}{"loginname": "mrdulin"}
		testHttpClient.AssertExpectations(t)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %+v, want: %+v", got, want)
		}

	})
}
