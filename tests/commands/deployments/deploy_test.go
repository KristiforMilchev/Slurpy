package deployments_test

import (
	"testing"

	"github.com/ovechkin-dm/mockio/mock"

	"slurpy/commands/deployments"
	"slurpy/implementations"
	"slurpy/tests/mocks"
)

var deployCommand deployments.DeployCommand

func TestF_Should_Deploy_Valid_Schema_On_Network(t *testing.T) {
	mock.SetUp(t)

	mockDeploymentService := mocks.MockDeploymentService{}
	deploymentService := mockDeploymentService.Init()
	deployCommand = deployments.DeployCommand{
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
			t.Log("Should not panic when calling deploy command with valid arguments")
			t.Fail()
		}
	}()

	path := "./../../test_data/deployment_schema.json"
	key := "group_name"
	network := "local"
	deployCommand.Execute(&path, &key, &network)
}
