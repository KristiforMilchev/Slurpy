package deployments_test

import (
	"testing"

	"github.com/ovechkin-dm/mockio/mock"

	"slurpy/commands/deployments"
	"slurpy/implementations"
	"slurpy/tests/mocks"
)

func TestF_Should_Retrive_A_List_Of_Deployments_With_Valid_Deployment_Group_Key(t *testing.T) {

	mock.SetUp(t)

	mockDeploymentService := mocks.MockDeploymentService{}
	deploymentService := mockDeploymentService.Init()
	allDeploymentsForKeyCommand := deployments.GetDeploymentCommand{
		Locator: implementations.Locator{
			NetworkService:    networkService.Init(),
			Storage:           storageService.Init(),
			DeploymentService: deploymentService,
			WalletService:     walletService.Init(),
			RpcService:        rpcService.Init(),
		},
	}

	defer func() {
		if r := recover(); r != nil {
			t.Log("Should not panic when calling get deployments with valid key")
			t.Fail()
		}
	}()

	key := "test_123"
	allDeploymentsForKeyCommand.Execute(&key)
}

func TestF_Should_Fail_To_Retrive_A_List_Of_Deployments_Withouth_A_Valid_Deployment_Group_Key(t *testing.T) {

	mock.SetUp(t)

	mockDeploymentService := mocks.MockDeploymentService{}
	deploymentService := mockDeploymentService.Init()
	allDeploymentsForKeyCommand := deployments.GetDeploymentCommand{
		Locator: implementations.Locator{
			NetworkService:    networkService.Init(),
			Storage:           storageService.Init(),
			DeploymentService: deploymentService,
			WalletService:     walletService.Init(),
			RpcService:        rpcService.Init(),
		},
	}

	defer func() {
		if r := recover(); r == nil {
			t.Log("Should panic when calling get deployments with invalid key command")
			t.Fail()
		}
	}()

	key := "test_12344444"
	allDeploymentsForKeyCommand.Execute(&key)
}
