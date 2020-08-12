package mocks

import (
	"github.com/mrdulin/graphql-go-cnode/utils"
	"github.com/stretchr/testify/mock"
)

type MockedHttpClient struct {
	mock.Mock
}

func (m *MockedHttpClient) Get(url string) (interface{}, error) {
	args := m.Called(url)
	return args.Get(0).(interface{}), args.Error(1)
}

func (m *MockedHttpClient) Post(url string, body interface{}) (interface{}, error) {
	args := m.Called(url, body)
	return args.Get(0).(interface{}), args.Error(1)
}

func (m *MockedHttpClient) HandleAPIError(res utils.ResponseMap) error {
	args := m.Called(res)
	return args.Error(0)
}
