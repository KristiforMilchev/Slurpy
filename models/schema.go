package models

type Schema struct {
	Contracts map[string]ContractConfig `json:"contracts"`
}
