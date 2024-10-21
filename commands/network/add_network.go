package network

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"

	"slurpy/implementations"
)

type AddNetwork struct {
	Locator implementations.Locator
}

func (a *AddNetwork) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "add-network [rpc] [network id] [name]",
		Short: "Adds a network to the database",
		Args:  cobra.MaximumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 3 {
				log.Fatal("Missing parameter, please add the required parameters")
			}

			rpc := args[0]
			networkId, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Error:", err)
			}
			networkName := args[2]

			a.Execute(&rpc, &networkId, &networkName)
		},
	}
}

func (a *AddNetwork) Execute(rpc *string, networkId *int, networkName *string) {
	a.Locator.Storage.Open()
	defer a.Locator.Storage.Close()

	_, err := a.Locator.NetworkService.Get(networkName)
	if err == nil {
		log.Panic("Network already exists with the same name!")
	}

	err = a.Locator.NetworkService.Add(rpc, networkId, networkName)

	if err != nil {
		log.Panic("Failed to save network to the database!")
	}
}
