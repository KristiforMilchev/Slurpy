package deployments

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"slurpy/implementations"
)

type AllDeploymentsCommand struct {
	Locator implementations.Locator
}

func (c *AllDeploymentsCommand) Executable() *cobra.Command {

	return &cobra.Command{
		Use:   "deployments",
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

	printer := implementations.TablePrinter{}

	for _, deployment := range deployments {
		dep := table.Row{deployment.Id, deployment.Contract, deployment.Date, deployment.Group}
		printer.Print([]table.Row{dep}, table.Row{
			"ID",
			"CONTRACT",
			"GROUP",
			"DATE",
		})

		var parameters []table.Row
		for _, parameter := range deployment.Options {
			param := table.Row{
				parameter,
			}

			parameters = append(parameters, param)
		}

		if len(parameters) > 0 {
			printer.Print(parameters, table.Row{"Value"})
		}
	}

	if len(deployments) == 0 {
		fmt.Println("No deployments found")
	}
}
