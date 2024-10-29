package implementations_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"slurpy/interfaces"
	"slurpy/models"
	"slurpy/tests/mocks"
)

var storage interfaces.Storage
var schema models.Schema
var bindTransOps bind.TransactOpts
var client ethclient.Client

func TestMain(m *testing.M) {

	// Run tests
	code := m.Run()
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
	os.Exit(code)
}
