package implementations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"slurpy/interfaces"
	"slurpy/models"
)

type DeploymentService struct {
	Storage interfaces.Storage
}

func (d *DeploymentService) GetDeployments() ([]models.Deployment, error) {
	d.Storage.Open()
	defer d.Storage.Close()

	sql := `
		SELECT d.id, d.contract, d.created_at, d.group_name
		FROM deployments AS d
	`
	rows, err := d.Storage.Query(&sql, &[]interface{}{})
	if err != nil {
		fmt.Println("Failed to fetch deployments")
		return nil, err
	}
	defer rows.Close()

	var data []models.Deployment

	for rows.Next() {
		var id int
		var contract, createdAt, group string
		err := rows.Scan(&id, &contract, &createdAt, &group)
		if err != nil {
			return nil, err
		}
		date, _ := time.Parse("2006-01-02 15:04:05", createdAt)

		data = append(data, models.Deployment{
			Id:       id,
			Contract: contract,
			Date:     date,
		})
	}

	return data, nil
}

func (d *DeploymentService) GetDeploymentByKey(key string) ([]models.Deployment, error) {
	d.Storage.Open()
	defer d.Storage.Close()

	sql := `
		SELECT d.id, d.contract, d.created_at, d.group_name
		FROM deployments AS d
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

		date, _ := time.Parse("2006-01-02 15:04:05", createdAt)
		deployment := &models.Deployment{
			Id:       id,
			Contract: contract,
			Date:     date,
			Options:  []string{},
		}
		deploymentMap[id] = deployment
		data = append(data, *deployment)

	}

	return data, nil
}
func (d *DeploymentService) DeployContracts(schema models.Schema, key *string, auth *bind.TransactOpts, client *ethclient.Client) error {
	// Store contract addresses by name; interface{} allows different types
	addresses := make(map[string]interface{})

	for contractName, config := range schema.Contracts {
		fmt.Println("Deploying deployment ", contractName)
		abiData, err := json.Marshal(config.Abi)
		if err != nil {
			log.Fatalf("Failed to marshal ABI: %v", err)
		}

		contractAbi, err := abi.JSON(bytes.NewReader(abiData))
		if err != nil {
			log.Fatalf("Failed to parse ABI for contract %s: %v", contractName, err)
			return err
		}

		bytecode := common.FromHex(config.Bytecode)

		var params []interface{}

		for _, dep := range config.Dependencies {
			switch dep.Type {
			case "deployment":
				address, exists := addresses[dep.Value]
				if !exists {
					log.Fatalf("Failed to find dependency %s for contract %s", contractName, contractName)
					return fmt.Errorf("dependency %s not found", contractName)
				}
				params = append(params, address)
			case "address":
				params = append(params, common.HexToAddress(dep.Value))
			case "int":
				val, err := strconv.Atoi(dep.Value)
				if err != nil {
					log.Fatal("Failed to parse value, type missmatch expected int got ", dep.Value)
				}
				params = append(params, val)
			case "float64":
				i, err := strconv.ParseInt(dep.Value, 10, 64)
				if err != nil {
					log.Fatal("Failed to parse value, type missmatch expected int got ", dep.Value)
				}
				params = append(params, i)
			case "bigInt":
				i, err := strconv.ParseInt(dep.Value, 10, 64)
				if err != nil {
					log.Fatal("Failed to parse value, type missmatch expected int got ", dep.Value)
				}
				val := big.NewInt(i)
				params = append(params, val)

			}

		}

		address, err := d.deploy(key, auth, contractAbi, &bytecode, client, &params)
		if err != nil {
			fmt.Println("Failed deployment")
			return err
		}
		addresses[contractName] = common.HexToAddress(address.Hex())
	}

	return nil
}

func (d *DeploymentService) deploy(key *string, auth *bind.TransactOpts, abi abi.ABI, bytecode *[]byte, client *ethclient.Client, params *[]interface{}) (common.Address, error) {

	address, tx, _, err := bind.DeployContract(auth, abi, *bytecode, client, *params...)
	if err != nil {
		fmt.Println(*params...)
		fmt.Println(err)
		return common.Address{}, err
	}

	fmt.Printf("Contract deployed at address: %s\n", address.Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

	d.Storage.Open()
	defer d.Storage.Close()

	insertDeploymentSQL := `
		INSERT INTO deployments (contract, created_at, group_name)
		VALUES ($1, datetime('now', 'localtime'), $2)
		RETURNING id
	`
	row := d.Storage.QuerySingle(&insertDeploymentSQL, &[]interface{}{
		address.Hex(),
		&key,
	})

	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Fatalf("Failed to save deployment details: %v", err)
		return common.Address{}, err
	}

	// insertParamsSQL := `
	// 	INSERT INTO deployment_parameters (parameter, deploymentId)
	// 	VALUES ($1, $2)
	// `
	// for _, param := range *params {
	// 	err := d.Storage.Exec(&insertParamsSQL, &[]interface{}{
	// 		&param,
	// 		&id,
	// 	})
	// 	if err != nil {
	// 		log.Fatalf("Failed to insert deployment parameters: %v", err)
	// 		return common.Address{}, err
	// 	}
	// }

	return address, nil
}
