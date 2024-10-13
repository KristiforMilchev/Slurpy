package interfaces

import "crypto/ecdsa"

type WalletService interface {
	Init(keys *[]string) error
	First() (*ecdsa.PrivateKey, error)
	WalletAt(index int) (*ecdsa.PrivateKey, error)
}
