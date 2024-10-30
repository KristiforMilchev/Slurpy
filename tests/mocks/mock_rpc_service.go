package mocks

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ovechkin-dm/mockio/mock"

	"slurpy/interfaces"
	"slurpy/tests"
)

type MockRpcService struct {
}

func (m *MockRpcService) Init() interfaces.RpcService {
	service := mock.Mock[interfaces.RpcService]()
	tests.LoadEnv()
	rpc := os.Getenv("RPC")
	fmt.Println(rpc)

	client, _ := ethclient.Dial(rpc)
	mock.When(service.GetClient()).ThenReturn(client)
	mock.When(service.SetClient(mock.Any[*string]())).ThenReturn(true, nil)
	return service
}
