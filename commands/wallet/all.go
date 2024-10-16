package wallet

import (
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"Slurpy/implementations"
)

type GetAllWalletsCommand struct {
	Locator implementations.Locator
}

func (g *GetAllWalletsCommand) Executable() *cobra.Command {

	return &cobra.Command{
		Use:   "wallets [network]",
		Args:  cobra.ExactArgs(1),
		Short: "List all wallets in a network",
		Run: func(cmd *cobra.Command, args []string) {
			network := args[0]
			g.Execute(&network)
		},
	}
}

func (g *GetAllWalletsCommand) Execute(network *string) {
	wallets, err := g.Locator.WalletService.GetWalletsForNetwork(network)
	if err != nil {
		log.Fatal("Failed to get all wallets for the given network", err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Key"})

	for _, wallet := range wallets {
		t.AppendRow(table.Row{wallet.Id, wallet.Key})
	}

	t.Render()
}
