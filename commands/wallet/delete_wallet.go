package wallet

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"

	"slurpy/implementations"
)

type DeleteWalletCommand struct {
	Locator implementations.Locator
}

func (d *DeleteWalletCommand) Executable() *cobra.Command {

	return &cobra.Command{
		Use:   "wallet delete [id]",
		Short: "Deletes a wallet based on it's database id",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {

			id, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal("Bad command argument, expected int received", args[0])
			}
			d.Execute(&id)
		},
	}
}

func (d *DeleteWalletCommand) Execute(id *int) {
	err := d.Locator.WalletService.DeleteWallet(id)

	if err != nil {
		log.Fatal(err, "Failed to delete wallet")
	}

	fmt.Println("Wallet has been removed from database!")
}
