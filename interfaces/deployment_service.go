package interfaces

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"slurpy/models"
)

type DeploymentService interface {
	GetDeployments() ([]models.Deployment, error)
	GetDeploymentByKey(key string) ([]models.Deployment, error)
	DeployContracts(schema models.Schema, key *string, auth *bind.TransactOpts, client *ethclient.Client) error
}
