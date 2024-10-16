package commands

import (
	"fmt"
	"log"

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
			g.Execute(&id)
		},
	}
}

func (g *GetCommand) Execute(id *string) {
	deployment, err := g.Locator.DeploymentService.GetDeploymentByKey(*id)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Failed to retrive deployment by id")
	}

	fmt.Println(deployment)
}
