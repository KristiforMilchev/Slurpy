package main

import (
	"log"
	"os"

	"slurpy/implementations"
	"slurpy/interfaces"
)

func SetupDatabase(config interfaces.Configuration) interfaces.Storage {
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
