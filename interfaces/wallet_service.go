package interfaces

import "crypto/ecdsa"

type WalletService interface {
	First(network *string) (*ecdsa.PrivateKey, error)
	WalletAt(index int, network *string) (*ecdsa.PrivateKey, error)
	AddWallet(key *string, network *string) error
}
