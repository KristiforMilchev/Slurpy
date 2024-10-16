package implementations

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"

	"Slurpy/interfaces"
	"Slurpy/models"
)

type WalletService struct {
	Storage interfaces.Storage
}

func (w *WalletService) First(network *string) (*ecdsa.PrivateKey, error) {
	sql := `
		SELECT wallet_key FROM wallets
		WHERE network = $1
		ORDER BY id
		LIMIT 1 OFFSET $2;
	`
	w.Storage.Open()
	defer w.Storage.Close()

	row := w.Storage.QuerySingle(&sql, &[]interface{}{
		&network,
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

func (w *WalletService) WalletAt(index int, network *string) (*ecdsa.PrivateKey, error) {
	sql := `
		SELECT wallet_key FROM wallets
		WHERE network = $1
		ORDER BY id
		LIMIT 1 OFFSET $2;
	`
	w.Storage.Open()
	defer w.Storage.Close()

	row := w.Storage.QuerySingle(&sql, &[]interface{}{
		&network,
		&index,
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

func (w *WalletService) DeleteWallet(id *int) error {
	sql := `
		DELETE FROM wallets
		WHERE id = $1
	`

	w.Storage.Open()
	defer w.Storage.Close()

	return w.Storage.Exec(&sql, &[]interface{}{
		&id,
	})
}

func (w *WalletService) GetWalletsForNetwork(network *string) ([]models.Wallet, error) {
	sql := `
		SELECT * FROM wallets
		WHERE network = $1
	`

	w.Storage.Open()
	defer w.Storage.Close()

	rows, err := w.Storage.Query(&sql, &[]interface{}{
		&network,
	})

	if err != nil {
		return []models.Wallet{}, err
	}

	var wallets []models.Wallet

	for rows.Next() {
		var wallet models.Wallet
		err := rows.Scan(&wallet.Id, &wallet.Key, &wallet.Network)
		if err != nil {
			return []models.Wallet{}, err
		}
		wallets = append(wallets, wallet)
	}

	return wallets, nil
}
