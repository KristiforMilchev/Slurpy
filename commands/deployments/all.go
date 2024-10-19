package deployments

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"slurpy/implementations"
)

type AllDeploymentsCommand struct {
	Locator implementations.Locator
}

func (c *AllDeploymentsCommand) Executable() *cobra.Command {

	return &cobra.Command{
		Use:   "all",
		Short: "List all deployments",
		Run: func(cmd *cobra.Command, args []string) {
			c.Execute()
		},
	}
}

func (c *AllDeploymentsCommand) Execute() {
	deployments, err := c.Locator.DeploymentService.GetDeployments()
	if err != nil {
		fmt.Println("Failed to get deployments, received an error!")
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Contract", "Date"})

	for _, deployment := range deployments {
		t.AppendRow(table.Row{deployment.Id, deployment.Contract, deployment.Date})
		fmt.Println()
		fmt.Println(deployment.Date)

	}

	t.Render()

	if len(deployments) == 0 {
		fmt.Println("No deployments found")
	}
}
