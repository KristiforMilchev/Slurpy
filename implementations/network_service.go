package implementations

import (
	"fmt"

	"Slurpy/interfaces"
	"Slurpy/models"
)

type NetworkService struct {
	Storage interfaces.Storage
}

func (n *NetworkService) Add(rpc *string, port *int, name *string) error {

	sql := `
		INSERT INTO networks
		(rpc, network_id, network_name)
		VALUES ($1,$2,$3)
		RETURNING 1
	`

	row := n.Storage.QuerySingle(&sql, &[]interface{}{
		&rpc,
		&port,
		&name,
	})

	var result int
	err := row.Scan(&result)

	if err != nil {
		fmt.Println("Failed to save query to database")
		return err
	}

	return nil
}

func (n *NetworkService) Get(name *string) (models.Network, error) {
	sql := `
		SELECT * FROM networks
		WHERE network_name = $1
	`

	row := n.Storage.QuerySingle(&sql, &[]interface{}{
		&name,
	})

	var network models.Network
	err := row.Scan(&network.Rpc, &network.NetworkId, &network.Name)
	if err != nil {
		fmt.Println("Failed to retrive network from database with name ", *name)
		return models.Network{}, err
	}

	return network, nil
}

func (n *NetworkService) Remove(name *string) error {
	sql := `
		DELETE FROM networks
		WHERE network_name = $1
	`

	row := n.Storage.QuerySingle(&sql, &[]interface{}{
		&name,
	})

	err := row.Scan()
	if err != nil {
		fmt.Println("Failed to delete network with name", name)
		return err
	}

	return nil
}
