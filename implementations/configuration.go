package implementations

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	values map[string]interface{}
	File   *string
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
