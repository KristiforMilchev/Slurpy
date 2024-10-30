package main

import (
	"slurpy/implementations"
	"slurpy/interfaces"
	"slurpy/repositories"
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
			DeploymentRepositoy: &repositories.DeploymentRepository{
				Storage: storage,
			},
		},
		NetworkService: &implementations.NetworkService{
			Storage: storage,
		},
	}
}
