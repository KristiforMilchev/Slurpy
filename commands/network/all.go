package network

import (
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"slurpy/implementations"
)

type GetAllNetworksCommand struct {
	Locator implementations.Locator
}

func (g *GetAllNetworksCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "networks",
		Short: "Gets a list of saved networks",
		Run: func(cmd *cobra.Command, args []string) {

			g.Execute()
		},
	}
}

func (g *GetAllNetworksCommand) Execute() {

	networks, err := g.Locator.NetworkService.All()

	if err != nil {
		log.Fatal("Failed to fetch networks from the database")
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Key"})

	for _, network := range networks {
		t.AppendRow(table.Row{network.Rpc, network.Name})
	}

	t.Render()

}
