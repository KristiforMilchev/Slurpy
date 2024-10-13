package implementations

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type RpcService struct {
	client *ethclient.Client
}

func (r *RpcService) GetClient() *ethclient.Client {
	return r.client
}

func (r *RpcService) SetClient(rpc *string) (bool, error) {
	var err error
	r.client, err = ethclient.Dial(*rpc)
	if err != nil {
		fmt.Printf("failed to connect to the Ethereum client: %v", err)
		return false, err
	}
	return true, nil
}
