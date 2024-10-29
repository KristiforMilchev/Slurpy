package implementations_test

import (
	"context"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ovechkin-dm/mockio/mock"

	"slurpy/implementations"
	"slurpy/interfaces"
	"slurpy/tests/mocks"
)

var deployementService interfaces.DeploymentService

func TestF_Should_Initialiaze(t *testing.T) {
	mock.SetUp(t)
	deployementService = &implementations.DeploymentService{
		Storage: storage,
	}
	local := "local"
	walletService := mocks.MockWalletService{}
	mockRpcService := mocks.MockRpcService{}
	rpcService := mockRpcService.Init()
	wallet, _ := walletService.Init().First(&local)
	client = *rpcService.GetClient()

	chainid, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("Failed to retrive chain ID from RPC")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(wallet, chainid)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	bindTransOps = *auth
	if deployementService == nil {
		t.Log("Should not fail to initialize deployment service")
		t.Fail()
	}
}
func TestF_Should_Deploy_A_Migration_With_Valid_Syntax(t *testing.T) {
	mock.SetUp(t)
	deploymentGroup := "valid_migration"

	err := deployementService.DeployContracts(schema, &deploymentGroup, &bindTransOps, &client)

	if err != nil {
		t.Log("Should not fail to deploy a valid migration scheme")
		t.Fail()
	}
}
