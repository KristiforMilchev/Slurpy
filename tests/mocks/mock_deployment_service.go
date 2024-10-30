package mocks

import (
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ovechkin-dm/mockio/mock"

	"slurpy/interfaces"
	"slurpy/models"
)

type MockDeploymentService struct {
}

func (m *MockDeploymentService) Init() interfaces.DeploymentService {
	service := mock.Mock[interfaces.DeploymentService]()
	mock.When(service.DeployContracts(mock.Any[models.Schema](), mock.Any[*string](), mock.Any[*bind.TransactOpts](), mock.Any[*ethclient.Client]())).ThenReturn(nil)
	mock.When(service.GetDeploymentByKey("test_123")).ThenReturn([]models.Deployment{
		{
			Name:     "Pk Swap",
			Contract: "0xtest1",
			Group:    "Initial Deployment",
			Date:     time.Now(),
			Options: []string{
				"123",
				"123",
				"123",
			},
			Id: 1,
		},
		{
			Name:     "Pk Swap 2",
			Contract: "0xtest2",
			Group:    "Initial Deployment",
			Date:     time.Now(),
			Options: []string{
				"123",
				"123",
				"123",
			},
			Id: 1,
		},
	}, nil)
	mock.When(service.GetDeploymentByKey("test_12344444")).ThenReturn([]models.Deployment{}, errors.New("failed to fetch deployments"))
	mock.When(service.GetDeployments()).ThenReturn([]models.Deployment{
		{
			Name:     "Pk Swap",
			Contract: "0xtest1",
			Group:    "Initial Deployment",
			Date:     time.Now(),
			Options: []string{
				"123",
				"123",
				"123",
			},
			Id: 1,
		},
		{
			Name:     "Pk Swap 2",
			Contract: "0xtest2",
			Group:    "Initial Deployment",
			Date:     time.Now(),
			Options: []string{
				"123",
				"123",
				"123",
			},
			Id: 1,
		},

		{
			Name:     "ERC 20",
			Contract: "0xtest3",
			Group:    "Second Deployment",
			Date:     time.Now(),
			Options: []string{
				"123",
				"123",
				"123",
			},
			Id: 1,
		},
	}, nil)

	return service
}
