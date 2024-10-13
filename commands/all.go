package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"Slurpy/implementations"
)

type AllCommand struct {
	Locator implementations.Locator
}

func (c *AllCommand) Executable() *cobra.Command {

	return &cobra.Command{
		Use:   "all",
		Short: "List all deployments",
		Run: func(cmd *cobra.Command, args []string) {
			deployments, err := c.Locator.DeploymentService.GetDeployments()
			if err != nil {
				fmt.Println("Failed to get deployments, received an error!")
				fmt.Println(err)
			}

			for _, deployment := range deployments {
				fmt.Println(deployment.Contract)
				fmt.Println(deployment.Date)
			}

			if len(deployments) == 0 {
				fmt.Println("No deployments found")
			}
		},
	}
}
