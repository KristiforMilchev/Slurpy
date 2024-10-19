package main

import (
	"slurpy/implementations"
	"slurpy/interfaces"
)

func Locator(storage interfaces.Storage) implementations.Locator {
	return implementations.Locator{
		WalletService: &implementations.WalletService{
			Storage: storage,
		},
		EncoderService: &implementations.EncoderService{},
		RpcService:     &implementations.RpcService{},
		Storage:        storage,
		DeploymentService: &implementations.DeploymentService{
			Storage: storage,
		},
		NetworkService: &implementations.NetworkService{
			Storage: storage,
		},
	}
}
