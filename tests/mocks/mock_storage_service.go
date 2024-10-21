package mocks

import (
	"github.com/ovechkin-dm/mockio/mock"

	"slurpy/interfaces"
)

type MockStorageService struct {
}

func (m *MockStorageService) Init() interfaces.Storage {
	service := mock.Mock[interfaces.Storage]()
	anyParams := mock.Any[[]interface{}]()
	anyQuery := mock.Any[string]()
	mock.When(service.Query(&anyQuery, &anyParams)).ThenReturn(true)
	mock.When(service.Exec(&anyQuery, &anyParams)).ThenReturn(true)
	mock.When(service.QuerySingle(&anyQuery, &anyParams)).ThenReturn(true)
	mock.When(service.Open()).ThenReturn(true)
	mock.When(service.Close()).ThenReturn(true)

	return service
}
