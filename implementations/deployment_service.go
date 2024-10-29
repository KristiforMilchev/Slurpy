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
		SELECT d.id, d.name, d.contract, d.created_at, d.group_name
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
		var name, contract, createdAt, group string
		err := rows.Scan(&id, &name, &contract, &createdAt, &group)
		if err != nil {
			return nil, err
		}
		date, _ := time.Parse("2006-01-02 15:04:05", createdAt)

		deployment := &models.Deployment{
			Id:       id,
			Name:     name,
			Contract: contract,
			Date:     date,
			Group:    group,
			Options:  []string{},
		}

		sqlParameters := `
			SELECT parameter FROM deployment_parameters
			WHERE deploymentId = $1
		`

		rows, err := d.Storage.Query(&sqlParameters, &[]interface{}{
			&id,
		})

		if err != nil {
			continue
		}
		var parameters []string
		for rows.Next() {
			var param string
			err := rows.Scan(&param)

			if err != nil {
				fmt.Println("Can't retrive parameter for deloyment with contract ", contract)
			}

			parameters = append(parameters, param)
		}
		deployment.Options = parameters
		data = append(data, *deployment)
	}

	return data, nil
}

func (d *DeploymentService) GetDeploymentByKey(key string) ([]models.Deployment, error) {
	d.Storage.Open()
	defer d.Storage.Close()

	sql := `
		SELECT d.id, d.name, d.contract, d.created_at, d.group_name
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

	var data []models.Deployment
	for rows.Next() {
		var id int
		var name, contract, createdAt, group string
		err := rows.Scan(&id, &name, &contract, &createdAt, &group)
		if err != nil {
			return nil, err
		}

		date, _ := time.Parse("2006-01-02 15:04:05", createdAt)
		deployment := &models.Deployment{
			Id:       id,
			Name:     name,
			Contract: contract,
			Date:     date,
			Group:    group,
			Options:  []string{},
		}

		sqlParameters := `
			SELECT parameter FROM deployment_parameters
			WHERE deploymentId = $1
		`

		rows, err := d.Storage.Query(&sqlParameters, &[]interface{}{
			&id,
		})

		if err != nil {
			continue
		}
		var parameters []string
		for rows.Next() {
			var param string
			err := rows.Scan(&param)

			if err != nil {
				fmt.Println("Can't retrive parameter for deloyment with contract ", contract)
			}

			parameters = append(parameters, param)
		}
		deployment.Options = parameters
		data = append(data, *deployment)
	}

	return data, nil
}
func (d *DeploymentService) DeployContracts(schema models.Schema, key *string, auth *bind.TransactOpts, client *ethclient.Client) error {
	addresses := make(map[string]interface{})

	for contractName, config := range schema.Contracts {
		abiData, err := json.Marshal(config.Abi)
		if err != nil {
			return err
		}

		contractAbi, err := abi.JSON(bytes.NewReader(abiData))
		if err != nil {
			return err
		}

		bytecode := common.FromHex(config.Bytecode)
		var beforeConverting []interface{}
		var params []interface{}

		for _, dep := range config.Dependencies {
			beforeConverting = append(beforeConverting, dep.Value)
			switch dep.Type {
			case "deployment":
				address, exists := addresses[dep.Value]
				if !exists {
					return fmt.Errorf("dependency %s not found", contractName)
				}
				params = append(params, address)
			case "address":
				params = append(params, common.HexToAddress(dep.Value))
			case "string":
				params = append(params, dep.Value)
			case "int":
				val, err := strconv.Atoi(dep.Value)
				if err != nil {
					return err
				}
				params = append(params, val)
			case "float64":
				i, err := strconv.ParseInt(dep.Value, 10, 64)
				if err != nil {
					return err
				}
				params = append(params, i)
			case "uint8":
				val, err := strconv.ParseUint(dep.Value, 10, 8)
				if err != nil {
					return err
				}
				params = append(params, uint8(val))
			case "uint32":
				val, err := strconv.ParseUint(dep.Value, 10, 32)
				if err != nil {
					return err
				}
				params = append(params, uint32(val))
			case "bigInt":
				i, ok := new(big.Int).SetString(dep.Value, 10)
				if !ok {
					return fmt.Errorf("invalid bigInt value: %s", dep.Value)
				}
				params = append(params, i)
			}

		}

		address, id, err := d.deploy(&config.Name, key, auth, contractAbi, &bytecode, client, &params)
		if err != nil {
			return err
		}

		err = d.SaveParameters(&beforeConverting, id)
		if err != nil {
			return err
		}

		addresses[contractName] = common.HexToAddress(address.Hex())
	}

	return nil
}

func (d *DeploymentService) deploy(name *string, key *string, auth *bind.TransactOpts, abi abi.ABI, bytecode *[]byte, client *ethclient.Client, params *[]interface{}) (common.Address, int, error) {
	address, tx, _, err := bind.DeployContract(auth, abi, *bytecode, client, *params...)
	if err != nil {
		fmt.Println(*params...)
		fmt.Println(err)
		return common.Address{}, 0, err
	}

	fmt.Printf("Contract deployed at address: %s\n", address.Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

	d.Storage.Open()
	defer d.Storage.Close()

	insertDeploymentSQL := `
		INSERT INTO deployments (contract, name, created_at, group_name)
		VALUES ($1, $2, datetime('now', 'localtime'), $3)
		RETURNING id
	`
	row := d.Storage.QuerySingle(&insertDeploymentSQL, &[]interface{}{
		address.Hex(),
		&name,
		&key,
	})

	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Fatalf("Failed to save deployment details: %v", err)
		return common.Address{}, 0, err
	}

	return address, id, nil
}

func (d *DeploymentService) SaveParameters(params *[]interface{}, id int) error {
	d.Storage.Open()
	defer d.Storage.Close()

	insertParamsSQL := `
		INSERT INTO deployment_parameters (parameter, deploymentId)
		VALUES ($1, $2)
	`

	for _, param := range *params {
		fmt.Println("saving ", param)
		str, ok := param.(string)
		if !ok {
			fmt.Print("Failed to convert parameter")
			continue
		}

		err := d.Storage.Exec(&insertParamsSQL, &[]interface{}{
			str,
			&id,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
