package implementations

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"

	"Slurpy/interfaces"
)

type WalletService struct {
	Storage interfaces.Storage
}

func (w *WalletService) First() (*ecdsa.PrivateKey, error) {
	sql := `
		SELECT wallet_key FROM wallets
		WHERE id = $1
	`
	w.Storage.Open()
	defer w.Storage.Close()

	row := w.Storage.QuerySingle(&sql, &[]interface{}{
		0,
	})

	var wallet string
	err := row.Scan(&wallet)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(wallet)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return privateKey, nil
}

func (w *WalletService) WalletAt(index int) (*ecdsa.PrivateKey, error) {
	sql := `
		SELECT wallet_key FROM wallets
		WHERE id = $1
	`
	w.Storage.Open()
	defer w.Storage.Close()

	row := w.Storage.QuerySingle(&sql, &[]interface{}{
		index,
	})

	var wallet string
	err := row.Scan(&wallet)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(wallet)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return privateKey, nil
}

func (w *WalletService) AddWallet(key *string, network *string) error {
	sql := `
		INSERT INTO wallets (wallet_key, network)
		VALUES ($1, $2)
		RETURNING id
	`
	w.Storage.Open()
	defer w.Storage.Close()

	row := w.Storage.QuerySingle(&sql, &[]interface{}{
		&key,
		&network,
	})

	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Fatalf("Failed to save deployment details: %v", err)
		return err
	}

	return nil
}
