package mocks

import (
	"github.com/ovechkin-dm/mockio/mock"

	"slurpy/interfaces"
)

type MockStorageService struct {
}

func (m *MockStorageService) Init() interfaces.Storage {

	service := mock.Mock[interfaces.Storage]()
	anyParams := mock.Any[*[]interface{}]()
	anyQuery := mock.Any[string]()
	mockRowsScanner := &MockRowsScanner{}
	mockRowScanner := &MockRowScanner{}
	mock.When(service.QuerySingle(&anyQuery, anyParams)).ThenReturn(mockRowScanner, nil)
	mock.When(service.Query(&anyQuery, anyParams)).ThenReturn(mockRowsScanner)
	mock.When(service.Exec(&anyQuery, anyParams)).ThenReturn(nil)
	mock.When(service.Open()).ThenReturn(true)
	mock.When(service.Close()).ThenReturn(true)

	return service
}
