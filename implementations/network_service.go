package implementations

import (
	"slurpy/interfaces"
	"slurpy/models"
)

type NetworkService struct {
	Storage interfaces.Storage
}

func (n *NetworkService) All() ([]models.Network, error) {
	n.Storage.Open()
	defer n.Storage.Close()
	sql := `
		SELECT * FROM networks
	`

	rows, err := n.Storage.Query(&sql, &[]interface{}{})

	if err != nil {
		return []models.Network{}, err
	}

	var networks []models.Network
	for rows.Next() {
		var network models.Network

		err := rows.Scan(&network.Name, &network.Rpc, &network.NetworkId)
		if err != nil {
			return []models.Network{}, err
		}

		networks = append(networks, network)
	}

	return networks, nil
}

func (n *NetworkService) Add(rpc *string, port *int, name *string) error {
	n.Storage.Open()
	defer n.Storage.Close()
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

	return err
}

func (n *NetworkService) Get(name *string) (models.Network, error) {
	n.Storage.Open()
	defer n.Storage.Close()
	sql := `
		SELECT * FROM networks
		WHERE network_name = $1
	`

	row := n.Storage.QuerySingle(&sql, &[]interface{}{
		&name,
	})

	var network models.Network
	err := row.Scan(&network.Name, &network.Rpc, &network.NetworkId)
	if err != nil {
		return models.Network{}, err
	}

	return network, nil
}

func (n *NetworkService) Remove(name *string) error {
	n.Storage.Open()
	defer n.Storage.Close()
	sql := `
		DELETE FROM networks
		WHERE network_name = $1
	`

	return n.Storage.Exec(&sql, &[]interface{}{
		&name,
	})
}
