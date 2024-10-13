package implementations

import (
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"Slurpy/interfaces"
	"Slurpy/models"
)

type DeploymentService struct {
	Storage interfaces.Storage
}

func (d *DeploymentService) GetDeployments() ([]models.Deployment, error) {
	d.Storage.Open()
	defer d.Storage.Close()

	sql := `
		SELECT d.id, d.contract, d.created_at, d.group_name, dp.parameter
		FROM deployments AS d
		LEFT JOIN deployment_parameters AS dp ON dp.deploymentId = d.id
	`
	rows, err := d.Storage.Query(&sql, &[]interface{}{})
	if err != nil {
		fmt.Println("Failed to fetch deployments")
		return nil, err
	}
	defer rows.Close()

	deploymentMap := make(map[int]*models.Deployment)
	var data []models.Deployment

	for rows.Next() {
		var id int
		var contract, createdAt, group, parameter string
		err := rows.Scan(&id, &contract, &createdAt, &group, &parameter)
		if err != nil {
			return nil, err
		}

		deployment, exists := deploymentMap[id]
		if !exists {
			date, _ := time.Parse("2006-01-02 15:04:05", createdAt)
			deployment = &models.Deployment{
				Id:       id,
				Contract: contract,
				Date:     date,
				Options:  []string{},
			}
			deploymentMap[id] = deployment
			data = append(data, *deployment)
		}

		if parameter != "" {
			deployment.Options = append(deployment.Options, parameter)
		}
	}

	return data, nil
}

func (d *DeploymentService) GetDeploymentByKey(key string) ([]models.Deployment, error) {
	d.Storage.Open()
	defer d.Storage.Close()

	sql := `
		SELECT d.id, d.contract, d.created_at, d.group_name, dp.parameter
		FROM deployments AS d
		LEFT JOIN deployment_parameters AS dp ON dp.deploymentId = d.id
		WHERE d.group_name = $1
	`
	rows, err := d.Storage.Query(&sql, &[]interface{}{
		&key,
	})

	if err != nil {
		fmt.Println("Failed to fetch deployments")
		return nil, err
	}
	defer rows.Close()

	deploymentMap := make(map[int]*models.Deployment)
	var data []models.Deployment

	for rows.Next() {
		var id int
		var contract, createdAt, group, parameter string
		err := rows.Scan(&id, &contract, &createdAt, &group, &parameter)
		if err != nil {
			return nil, err
		}

		deployment, exists := deploymentMap[id]
		if !exists {
			date, _ := time.Parse("2006-01-02 15:04:05", createdAt)
			deployment = &models.Deployment{
				Id:       id,
				Contract: contract,
				Date:     date,
				Options:  []string{},
			}
			deploymentMap[id] = deployment
			data = append(data, *deployment)
		}

		if parameter != "" {
			deployment.Options = append(deployment.Options, parameter)
		}
	}

	return data, nil
}

func (d *DeploymentService) Deploy(key *string, auth *bind.TransactOpts, abi *abi.ABI, bytecode *[]byte, client *ethclient.Client, params *[]interface{}) error {
	address, tx, _, err := bind.DeployContract(auth, *abi, *bytecode, client, *params...)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}

	fmt.Printf("Contract deployed at address: %s\n", address.Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

	d.Storage.Open()
	defer d.Storage.Close()

	insertDeploymentSQL := `
		INSERT INTO deployments (contract, created_at, group)
		VALUES ($1, datetime('now', 'localtime'), '')
		RETURNING id
	`
	row := d.Storage.QuerySingle(&insertDeploymentSQL, &[]interface{}{
		address.Hex(),
	})

	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Fatalf("Failed to save deployment details: %v", err)
		return err
	}

	insertParamsSQL := `
		INSERT INTO deployment_parameters (parameter, deploymentId)
		VALUES ($1, $2)
	`
	for _, param := range *params {
		err := d.Storage.Exec(&insertParamsSQL, &[]interface{}{
			&param,
			&id,
		})

		if err != nil {
			log.Fatalf("Failed to insert deployment parameters: %v", err)
			return err
		}
	}

	return nil
}
