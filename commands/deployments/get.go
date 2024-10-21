package deployments

import (
	"fmt"
	"log"

	"github.com/jedib0t/go-pretty/v6/table"
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
		Args:  cobra.ExactArgs(1),
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
	deployments, err := g.Locator.DeploymentService.GetDeploymentByKey(*id)
	if err != nil {
		log.Fatal("Failed to retrive deployment by id")
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
