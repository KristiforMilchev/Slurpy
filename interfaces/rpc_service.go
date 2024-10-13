package interfaces

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type RpcService interface {
	GetClient() *ethclient.Client
	SetClient(rpc *string) (bool, error)
}
