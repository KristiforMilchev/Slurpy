package deployments

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"

	"slurpy/implementations"
	"slurpy/models"
)

type DeployCommand struct {
	Locator implementations.Locator
}

func (d *DeployCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy [path] [network optional] [group optional]",
		Short: "Deploy a smart contract migration",
		Args:  cobra.MaximumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 3 {
				log.Fatalf("Parameters are required, expected 3 got %v", len(args))
			}
			path := args[0]
			network := args[1]
			key := ""
			if len(args) > 1 {
				key = args[1]
			}

			d.Execute(&path, &key, &network)
		},
	}
}

func (d *DeployCommand) Execute(path *string, key *string, network *string) {
	var schema models.Schema
	data, err := os.ReadFile(*path)
	if err != nil {
		log.Fatal("Failed to find file, may not exist!")
	}
	err = json.Unmarshal(data, &schema)
	if err != nil {
		log.Fatal("Failed to parse deployment file, please check for syntax errors!")
	}
	wallet, err := d.Locator.WalletService.WalletAt(0, network)
	if err != nil {
		log.Fatal("Failed to retrive wallet")
	}

	networkData, err := d.Locator.NetworkService.Get(network)
	if err != nil {
		log.Fatal("Failed to retrive netowrk by name")
	}

	fmt.Println(&networkData.Rpc)
	d.Locator.RpcService.SetClient(&networkData.Rpc)
	rpc := d.Locator.RpcService.GetClient()
	chainid, err := rpc.ChainID(context.Background())
	if err != nil {
		log.Fatal("Failed to retrive chain ID from RPC")
	}
	auth, err := bind.NewKeyedTransactorWithChainID(wallet, chainid)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}
	auth.GasLimit = uint64(6000000)
	err = d.Locator.DeploymentService.DeployContracts(schema, key, auth, rpc)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Failed to deploy migration!")
	}

	fmt.Println("Deployment migrated")
}
