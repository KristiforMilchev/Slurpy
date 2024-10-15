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

	configuration = &implementations.Configuration{}
	configuration.Load()

	storage := SetupDatabase(configuration)
	locator = Locator(storage)

	var rootCmd = &cobra.Command{Use: "slurpy"}

	all := commands.AllCommand{
		Locator: locator,
	}

	configCommand := commands.ConfigCommand{
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

	rpc := "HTTP://127.0.0.1:7545"
	networkId := 5777
	network := "local"

	addNetworkCommand.Execute(&rpc, &networkId, &network)

	privateKey := "bfc9f203f59ea2b5d664f0b6101df961dc834e6e8cbca6a1c364b9b57ae3f04f"

	addWalletCommand.Execute(&privateKey, &network)

	path := "deployment.json"
	key := "test"

	deployCommand.Execute(&path, &key, &network)

	rootCmd.AddCommand(all.Executable())
	rootCmd.AddCommand(configCommand.Executable())
	rootCmd.AddCommand(getCommand.Executable())
	rootCmd.AddCommand(deployCommand.Executable())
	rootCmd.AddCommand(addWalletCommand.Executable())
	rootCmd.AddCommand(addNetworkCommand.Executable())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
