package models

type ContractConfig struct {
	Bytecode     string        `json:"bytecode"`
	Abi          []interface{} `json:"abi"`
	Dependencies []interface{} `json:"dependencies"`
}