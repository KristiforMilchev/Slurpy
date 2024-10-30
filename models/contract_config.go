package models

type ContractConfig struct {
	Name         string        `json:"name"`
	Bytecode     string        `json:"bytecode"`
	Abi          []interface{} `json:"abi"`
	Dependencies []Dependency  `json:"dependencies"`
}
