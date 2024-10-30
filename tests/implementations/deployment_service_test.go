package implementations_test

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ovechkin-dm/mockio/mock"

	"slurpy/implementations"
	"slurpy/interfaces"
	"slurpy/tests/mocks"
)

var deployementService interfaces.DeploymentService

func setup() {
	storageMock := mocks.MockStorageService{}
	storage = storageMock.Init()
	file, err := os.ReadFile("../test_data/deployment_schema.json")
	if err != nil {
		log.Fatalf("Failed to retrive file, %v", err)
	}

	err = json.Unmarshal(file, &schema)
	if err != nil {
		log.Fatal("Failed to parse deployment file, please check for syntax errors!")
	}

	mockDeploymentRepository := mocks.MockDeploymentRepository{}

	deployementService = &implementations.DeploymentService{
		DeploymentRepositoy: mockDeploymentRepository.Init(),
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

}
func TestF_Should_Deploy_A_Migration_With_Valid_Syntax(t *testing.T) {
	mock.SetUp(t)
	setup()
	deploymentGroup := "valid_migration"

	err := deployementService.DeployContracts(schema, &deploymentGroup, &bindTransOps, &client)

	if err != nil {
		t.Log("Should not fail to deploy a valid migration scheme")
		t.Fail()
	}
}
