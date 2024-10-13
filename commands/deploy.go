package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"Slurpy/implementations"
)

type DeployCommand struct {
	Locator implementations.Locator
}

func (d *DeployCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy [bytecode(Hex)] [abi(string)] [group (string) optional]",
		Short: "Deploy a new contract",
		Args:  cobra.MaximumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			bytecodeHex := args[0]
			abi := args[1]
			key := ""
			if len(args) > 2 {
				key = args[2]
			}

			fmt.Printf("Deploying contract with bytecode: %s, abi: %s, key: %s...\n", bytecodeHex, abi, key)
		},
	}
}
