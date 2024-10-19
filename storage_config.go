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
	dbFile := "./db.sql"
	dbData := FileFromExecutable(&dbFile)
	tablesData, err := os.ReadFile(*dbData)

	if err != nil {
		log.Fatal("Couldn't find database file, aborting!")
	}

	tables := string(tablesData)
	storage.New(&dbPath, &dbName, &tables)
	storage.Open()
	storage.Initialize()
	return storage
}
