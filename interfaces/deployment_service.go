package interfaces

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"Slurpy/models"
)

type DeploymentService interface {
	GetDeployments() ([]models.Deployment, error)
	GetDeploymentByKey(key string) ([]models.Deployment, error)
	Deploy(key *string, auth *bind.TransactOpts, abi *abi.ABI, bytecode *[]byte, client *ethclient.Client, params *[]interface{}) error
}
