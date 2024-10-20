package mocks

import (
	"errors"

	"github.com/ovechkin-dm/mockio/mock"

	"slurpy/interfaces"
	"slurpy/models"
)

type MockNetwokService struct {
}

func (m *MockNetwokService) Init() interfaces.NetworkService {
	service := mock.Mock[interfaces.NetworkService]()
	mock.When(service.All()).ThenReturn(func(args []any) ([]models.Network, error) {
		var err error
		var networks []models.Network
		networks = append(networks, models.Network{
			Rpc:       "HTTPS://127.0.0.1:2245",
			NetworkId: 2222,
			Name:      "Local",
		})
		return networks, err
	})

	networkRpc := "HTTPS://127.0.0.1:2245"
	netwokId := 2222
	name := "local"
	mock.When(service.Add(&networkRpc, &netwokId, &name)).ThenReturn(nil)

	mock.When(service.Get(&name)).ThenReturn(func(args []any) (models.Network, error) {
		var err error
		network := models.Network{
			Rpc:       "HTTPS://127.0.0.1:2245",
			NetworkId: 2222,
			Name:      "Local",
		}
		return network, err
	})

	nonExistingNetowrk := "asd"
	mock.When(service.Get(&nonExistingNetowrk)).ThenReturn(func(args []any) (models.Network, error) {
		network := models.Network{}
		err := errors.New("no rows affected")
		return network, err
	})

	mock.When(service.Remove(&name)).ThenReturn(nil)
	mock.When(service.Remove(&nonExistingNetowrk)).ThenReturn(errors.New("no rows affected"))

	return service
}
