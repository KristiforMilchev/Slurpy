package network

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"slurpy/implementations"
)

type RemoveNetwork struct {
	Locator implementations.Locator
}

func (a *RemoveNetwork) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "remove-network [name]",
		Short: "Removes a network from the database and all it's dependencies",
		Args:  cobra.MaximumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatal("Missing parameter, please add the required parameters")
			}

			rpc := args[0]

			a.Execute(&rpc)
		},
	}
}

func (a *RemoveNetwork) Execute(name *string) {
	err := a.Locator.NetworkService.Remove(name)

	if err.Error() == "no rows affected" {
		panic("Network doesn't exist")
	}

	if err != nil {
		log.Panicf("Failed to remove a network with name: %v", *name)
	}

	fmt.Println("Network has been removed")
}
