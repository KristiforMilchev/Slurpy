package implementations

import (
	"crypto/ecdsa"
	"errors"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

type WalletService struct {
	wallets []ecdsa.PrivateKey
}

func (w *WalletService) Init(keys *[]string) error {
	for _, key := range *keys {
		privateKey, err := crypto.HexToECDSA(key)
		if err != nil {
			log.Fatalf("Failed to load private key: %v", err)
			return err
		}

		w.wallets = append(w.wallets, *privateKey)
	}
	return nil
}

func (w *WalletService) First() (*ecdsa.PrivateKey, error) {
	if len(w.wallets) == 0 {
		return nil, errors.New("collection is empty")
	}

	return &w.wallets[0], nil
}

func (w *WalletService) WalletAt(index int) (*ecdsa.PrivateKey, error) {
	if len(w.wallets) < index {
		return nil, errors.New("collection is empty")

	}

	return &w.wallets[index], nil
}
