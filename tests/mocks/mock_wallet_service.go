package mocks

import (
	"crypto/ecdsa"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/consensys/gnark-crypto/ecc/bls24-317/twistededwards/eddsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ovechkin-dm/mockio/mock"

	"slurpy/interfaces"
	"slurpy/models"
	"slurpy/tests"
)

type MockWalletService struct {
}

func (m *MockWalletService) Init() interfaces.WalletService {
	tests.LoadEnv()

	privateKeys := strings.Split(os.Getenv("PRIVATE_KEY"), ",")
	if len(privateKeys) == 0 {
		log.Fatalf("No private keys found in PRIVATE_KEY")
	}

	service := mock.Mock[interfaces.WalletService]()
	mock.When(service.AddWallet(mock.Any[*string](), mock.Any[*string]())).ThenReturn(nil)
	mock.When(service.DeleteWallet(mock.Any[*int]())).ThenReturn(nil)
	key, err := m.ImportPrivateKeyFromString(strings.TrimPrefix(privateKeys[0], "0x"))
	if err != nil {
		log.Fatalf("Failed to import private key: %v", err)
	}
	mock.When(service.First(mock.Any[*string]())).ThenReturn(key, nil)
	var keys []models.Wallet

	for i, key := range privateKeys {
		keys = append(keys, models.Wallet{
			Id:      i,
			Key:     key,
			Network: "local",
		})
	}
	local := "local"
	mock.When(service.GetWalletsForNetwork(mock.Any[*string]())).ThenReturn(keys, nil)
	mock.When(service.WalletAt(0, &local)).ThenReturn(key, nil)
	mock.When(service.WalletAt(1, &local)).ThenReturn(keys[1].Key, nil)
	mock.When(service.WalletAt(55, &local)).ThenReturn(eddsa.PrivateKey{}, errors.New("wallet at index doesn't exist"))

	return service
}

func (m *MockWalletService) ImportPrivateKeyFromString(hexKey string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	privateKey.PublicKey = *publicKeyECDSA

	return privateKey, nil
}
