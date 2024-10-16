package wallet

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"Slurpy/implementations"
)

type AddWallet struct {
	Locator implementations.Locator
}

func (a *AddWallet) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "add-wallet [private key] [network]",
		Short: "Adds a wallet for a given network",
		Args:  cobra.MaximumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			network := args[1]

			a.Execute(&path, &network)
		},
	}
}

func (a *AddWallet) Execute(privateKey *string, network *string) {
	a.Locator.Storage.Open()
	defer a.Locator.Storage.Close()

	_, err := a.Locator.NetworkService.Get(network)

	if err != nil {
		log.Fatal("Failed to find network with name", network)
		return
	}

	err = a.Locator.WalletService.AddWallet(privateKey, network)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Failed to add wallet to database, aborting!")
	}
}
