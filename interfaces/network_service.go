package interfaces

import "Slurpy/models"

type NetworkService interface {
	Add(rpc *string, port *int, name *string) error
	Get(name *string) (models.Network, error)
	Remove(name *string) error
}
