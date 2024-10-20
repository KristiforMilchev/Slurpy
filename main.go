package main

import (
	"os"

	"github.com/spf13/cobra"

	"slurpy/commands/deployments"
	"slurpy/commands/network"
	"slurpy/commands/wallet"
	"slurpy/implementations"
	"slurpy/interfaces"
)

var locator implementations.Locator
var configuration interfaces.Configuration

func main() {
	localSettingsFile := "./settings.json"
	settingsFile := FileFromExecutable(&localSettingsFile)
	configuration = &implementations.Configuration{
		File: settingsFile,
	}

	if exists := configuration.Exists(); !exists {
		configuration.Create()
	}

	configuration.Load()

	storage := SetupDatabase(configuration)
	locator = Locator(storage)

	var rootCmd = &cobra.Command{Use: "slurpy"}

	getAllDeployments := deployments.AllDeploymentsCommand{
		Locator: locator,
	}
	getDeploymentCommand := deployments.GetDeploymentCommand{
		Locator: locator,
	}
	deployCommand := deployments.DeployCommand{
		Locator: locator,
	}
	addWalletCommand := wallet.AddWallet{
		Locator: locator,
	}
	addNetworkCommand := network.AddNetwork{
		Locator: locator,
	}
	getNetworksCommand := network.GetAllNetworksCommand{
		Locator: locator,
	}
	getNetworkWallets := wallet.GetAllWalletsCommand{
		Locator: locator,
	}
	removeNetworkCommand := network.RemoveNetwork{
		Locator: locator,
	}
	deleteWallet := wallet.DeleteWalletCommand{
		Locator: locator,
	}

	rootCmd.AddCommand(getAllDeployments.Executable())
	rootCmd.AddCommand(getDeploymentCommand.Executable())
	rootCmd.AddCommand(deployCommand.Executable())
	rootCmd.AddCommand(addWalletCommand.Executable())
	rootCmd.AddCommand(getNetworksCommand.Executable())
	rootCmd.AddCommand(addNetworkCommand.Executable())
	rootCmd.AddCommand(getNetworkWallets.Executable())
	rootCmd.AddCommand(removeNetworkCommand.Executable())
	rootCmd.AddCommand(deleteWallet.Executable())

	if err := rootCmd.Execute(); err != nil {
		// fmt.Println(err)
		os.Exit(1)
	}
}
