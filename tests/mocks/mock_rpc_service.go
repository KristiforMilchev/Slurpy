package mocks

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/ovechkin-dm/mockio/mock"

	"slurpy/interfaces"
)

type MockRpcService struct {
}

func (m *MockRpcService) Init() interfaces.RpcService {
	service := mock.Mock[interfaces.RpcService]()
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	rpc := os.Getenv("RPC")

	client, _ := ethclient.Dial(rpc)
	mock.When(service.GetClient()).ThenReturn(client)
	mock.When(service.SetClient(mock.Any[*string]())).ThenReturn(true, nil)
	return service
}
