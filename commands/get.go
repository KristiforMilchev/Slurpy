package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"Slurpy/implementations"
)

type GetCommand struct {
	Locator implementations.Locator
}

func (g *GetCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "get [parameter id]",
		Short: "Get a deployment parameter by ID",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			param := args[1]
			// Fetch deployment parameter by id
			fmt.Printf("Fetching deployment parameter '%s' with id '%s'...\n", param, id)
			// Logic to get the parameter
		},
	}
}
