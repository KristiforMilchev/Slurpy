package mocks

import (
	"errors"
	"time"

	"github.com/ovechkin-dm/mockio/mock"

	"slurpy/interfaces"
	"slurpy/models"
)

type MockDeploymentRepository struct {
}

func (m *MockDeploymentRepository) Init() interfaces.DeploymentRepository {
	service := mock.Mock[interfaces.DeploymentRepository]()
	mock.When(service.GetAll()).ThenReturn([]models.Deployment{
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
	mock.When(service.GetDeploymentByKey("test_12344444")).ThenReturn([]models.Deployment{}, errors.New("failed to fetch deployments"))
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

	mock.When(service.SaveDeployment(mock.Any[*string](), mock.Any[*string](), mock.Any[*string]())).ThenReturn(1, nil)
	mock.When(service.SaveParameters(mock.Any[*[]interface{}](), mock.Any[int]())).ThenReturn(nil)

	return service
}
