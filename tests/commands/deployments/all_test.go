package deployments_test

import (
	"testing"

	"github.com/ovechkin-dm/mockio/mock"

	"slurpy/commands/deployments"
	"slurpy/implementations"
	"slurpy/tests/mocks"
)

func TestF_Should_Retrive_A_List_Of_Deployments(t *testing.T) {
	mock.SetUp(t)
	mockDeploymentService := mocks.MockDeploymentService{}
	deploymentService := mockDeploymentService.Init()
	allDeploymentsCommand := deployments.AllDeploymentsCommand{
		Locator: implementations.Locator{
			NetworkService:    networkService.Init(),
			Storage:           storageService.Init(),
			DeploymentService: deploymentService,
		},
	}

	defer func() {
		if r := recover(); r != nil {
			t.Log("Should not panic when calling all deployments command")
			t.Fail()
		}
	}()

	allDeploymentsCommand.Execute()
}
