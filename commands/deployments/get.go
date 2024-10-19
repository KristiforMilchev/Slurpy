package deployments

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"slurpy/implementations"
)

type GetDeploymentCommand struct {
	Locator implementations.Locator
}

func (g *GetDeploymentCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "deployment [id]",
		Short: "Get a deployment parameter by ID",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatal("deployment key is required")
			}

			id := args[0]
			g.Execute(&id)
		},
	}
}

func (g *GetDeploymentCommand) Execute(id *string) {
	deployment, err := g.Locator.DeploymentService.GetDeploymentByKey(*id)
	if err != nil {
		log.Fatal("Failed to retrive deployment by id")
	}

	fmt.Println(deployment)
}
