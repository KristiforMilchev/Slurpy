package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"Slurpy/commands"
	"Slurpy/commands/network"
	"Slurpy/commands/wallet"
	"Slurpy/implementations"
	"Slurpy/interfaces"
)

var locator implementations.Locator
var configuration interfaces.Configuration

func main() {
	settingsFile := "settings.json"
	configuration = &implementations.Configuration{
		File: &settingsFile,
	}
	configuration.Load()

	storage := SetupDatabase(configuration)
	locator = Locator(storage)

	var rootCmd = &cobra.Command{Use: "slurpy"}

	all := commands.AllCommand{
		Locator: locator,
	}

	getCommand := commands.GetCommand{
		Locator: locator,
	}
	deployCommand := commands.DeployCommand{
		Locator: locator,
	}
	addWalletCommand := wallet.AddWallet{
		Locator: locator,
	}
	addNetworkCommand := network.AddNetwork{
		Locator: locator,
	}

	getNetworkWallets := wallet.GetAllWalletsCommand{
		Locator: locator,
	}
	deleteWallet := wallet.DeleteWalletCommand{
		Locator: locator,
	}

	rootCmd.AddCommand(all.Executable())
	rootCmd.AddCommand(getCommand.Executable())
	rootCmd.AddCommand(deployCommand.Executable())
	rootCmd.AddCommand(addWalletCommand.Executable())
	rootCmd.AddCommand(addNetworkCommand.Executable())
	rootCmd.AddCommand(getNetworkWallets.Executable())
	rootCmd.AddCommand(deleteWallet.Executable())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
