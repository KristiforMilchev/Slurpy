package implementations_test

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"slurpy/interfaces"
	"slurpy/models"
)

var storage interfaces.Storage
var schema models.Schema
var bindTransOps bind.TransactOpts
var client ethclient.Client

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
