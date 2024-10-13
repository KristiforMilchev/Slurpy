package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"Slurpy/implementations"
)

type ConfigCommand struct {
	Locator implementations.Locator
}

func (c *ConfigCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "config [contracts - Path to folder] [wallets - Path to file]",
		Short: "Set the default configuration path",
		Run: func(cmd *cobra.Command, args []string) {
			contracts := args[0]
			wallets := ""
			fmt.Println("Updating Configuration paths")
			fmt.Println(contracts)
			fmt.Println(wallets)
		},
	}
}
