package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"Slurpy/commands"
	"Slurpy/implementations"
	"Slurpy/interfaces"
)

var locator implementations.Locator
var configuration interfaces.Configuration

func main() {

	configuration = &implementations.Configuration{}
	configuration.Load()

	storage := setupDatabase(configuration)

	locator = implementations.Locator{
		WalletService:  &implementations.WalletService{},
		EncoderService: &implementations.EncoderService{},
		RpcService:     &implementations.RpcService{},
		Storage:        storage,
		DeploymentService: &implementations.DeploymentService{
			Storage: storage,
		},
	}

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
	migrateCommand := commands.MigrateCommand{
		Locator: locator,
	}

	rootCmd.AddCommand(all.Executable())
	rootCmd.AddCommand(configCommand.Executable())
	rootCmd.AddCommand(getCommand.Executable())
	rootCmd.AddCommand(deployCommand.Executable())
	rootCmd.AddCommand(migrateCommand.Executable())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setupDatabase(config interfaces.Configuration) interfaces.Storage {
	storage := &implementations.Storage{}

	dbPath := config.GetKey("DbPath").(string)
	dbName := config.GetKey("DbName").(string)

	tablesData, err := os.ReadFile("db.sql")

	if err != nil {
		log.Fatal("Couldn't find database file, aborting!")
	}

	tables := string(tablesData)
	storage.New(&dbPath, &dbName, &tables)
	storage.Open()
	storage.Initialize()
	return storage
}
