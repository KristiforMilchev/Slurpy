package interfaces

import "slurpy/models"

type NetworkService interface {
	All() ([]models.Network, error)
	Add(rpc *string, port *int, name *string) error
	Get(name *string) (models.Network, error)
	Remove(name *string) error
}
