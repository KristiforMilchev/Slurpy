package wallet

import (
	"log"

	"github.com/spf13/cobra"

	"slurpy/implementations"
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
			if len(args) < 2 {
				log.Fatal("Missing arguments, please add a private key and a network")
			}

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
		log.Fatalf("Failed to find network with name %v", network)
		return
	}

	err = a.Locator.WalletService.AddWallet(privateKey, network)
	if err != nil {
		log.Fatal("Failed to save add wallet!")
	}
}
