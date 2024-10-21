package networks_test

import (
	"os"
	"testing"

	"slurpy/commands/network"
	"slurpy/tests/mocks"
)

var networkService mocks.MockNetwokService
var storageService mocks.MockStorageService
var removeNetworkCommand network.RemoveNetwork

func TestMain(m *testing.M) {
	// Run tests
	code := m.Run()
	os.Exit(code)
}
