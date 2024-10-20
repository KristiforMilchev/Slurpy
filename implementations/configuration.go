package implementations

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	values map[string]interface{}
	File   *string
}

func (c *Configuration) Exists() bool {
	_, err := os.Open(*c.File)
	return err == nil
}

func (c *Configuration) Create() bool {
	file, err := os.Create("./settings.json")
	if err != nil {
		panic("Failed to create default configuration file")
	}

	_, err = file.WriteString(`
	{
		"DbPath": "./data/",
		"DbName": "slurpy.db"
	}
	`)

	return err == nil
}

func (config *Configuration) Load() bool {
	config.values = make(map[string]interface{})

	file, err := os.Open(*config.File)
	if err != nil {
		panic("Configuration file doesn't exist.")
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config.values); err != nil {
		panic("Malformed Configuration file!")
	}

	defer file.Close()

	return true
}

func (config *Configuration) GetKey(name string) interface{} {
	result := config.values[name]
	return result
}
