package implementations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"slurpy/interfaces"
	"slurpy/models"
)

type DeploymentService struct {
	DeploymentRepositoy interfaces.DeploymentRepository
}

func (d *DeploymentService) GetDeployments() ([]models.Deployment, error) {
	return d.DeploymentRepositoy.GetAll()
}

func (d *DeploymentService) GetDeploymentByKey(key string) ([]models.Deployment, error) {
	return d.DeploymentRepositoy.GetDeploymentByKey(key)
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

		err = d.DeploymentRepositoy.SaveParameters(&beforeConverting, id)
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
	contractAddress := address.Hex()
	id, err := d.DeploymentRepositoy.SaveDeployment(&contractAddress, name, key)
	if err != nil {
		log.Fatalf("Failed to save deployment details: %v", err)
		return common.Address{}, 0, err
	}

	return address, id, nil
}
