package networks_test

import (
	"os"
	"testing"

	"github.com/ovechkin-dm/mockio/mock"

	"slurpy/commands/network"
	"slurpy/implementations"
	"slurpy/tests/mocks"
)

var networkService mocks.MockNetwokService
var removeNetworkCommand network.RemoveNetwork

func TestMain(m *testing.M) {
	// Run tests
	code := m.Run()
	os.Exit(code)
}

func TestF_Should_Remove_Existing_Network(t *testing.T) {
	mock.SetUp(t)
	m := networkService.Init()

	removeNetworkCommand = network.RemoveNetwork{
		Locator: implementations.Locator{
			NetworkService: m,
		},
	}
	defer func() {
		if r := recover(); r == nil {
			t.Log("Should not panic when removing existing network")
			t.Fail()
		}
	}()

	name := "local"
	removeNetworkCommand.Execute(&name)
}

func TestF_Should_Fail_To_Remove_Non_Existing_Network(t *testing.T) {
	mock.SetUp(t)
	m := networkService.Init()

	removeNetworkCommand = network.RemoveNetwork{
		Locator: implementations.Locator{
			NetworkService: m,
		},
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but function did not panic")

		}
	}()

	name := "asd"
	removeNetworkCommand.Execute(&name)
}
