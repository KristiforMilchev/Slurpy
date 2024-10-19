package interfaces

import (
	"crypto/ecdsa"

	"slurpy/models"
)

type WalletService interface {
	First(network *string) (*ecdsa.PrivateKey, error)
	WalletAt(index int, network *string) (*ecdsa.PrivateKey, error)
	AddWallet(key *string, network *string) error
	DeleteWallet(id *int) error
	GetWalletsForNetwork(network *string) ([]models.Wallet, error)
}
