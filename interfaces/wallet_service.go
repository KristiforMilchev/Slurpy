package interfaces

import "crypto/ecdsa"

type WalletService interface {
	First() (*ecdsa.PrivateKey, error)
	WalletAt(index int) (*ecdsa.PrivateKey, error)
	AddWallet(key *string) error
}
