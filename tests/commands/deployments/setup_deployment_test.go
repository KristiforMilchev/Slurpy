package deployments_test

import (
	"os"
	"testing"

	"slurpy/tests/mocks"
)

var networkService mocks.MockNetwokService
var storageService mocks.MockStorageService
var walletService mocks.MockWalletService
var rpcService mocks.MockRpcService

func TestMain(m *testing.M) {
	// Run tests
	code := m.Run()
	os.Exit(code)
}
