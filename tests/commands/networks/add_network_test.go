package networks_test

import (
	"testing"

	"github.com/ovechkin-dm/mockio/mock"

	"slurpy/commands/network"
	"slurpy/implementations"
)

func TestF_Add_Network_Command(t *testing.T) {
	mock.SetUp(t)
	m := networkService.Init()
	s := storageService.Init()
	addNetworkCommand := network.AddNetwork{
		Locator: implementations.Locator{
			NetworkService: m,
			Storage:        s,
		},
	}

	defer func() {
		if r := recover(); r != nil {
			t.Log("Should not panic when adding a network with valid data")
			t.Fail()
		}
	}()

	rpc := "HTTPS://127.0.0.1:2245"
	networkId := 2222
	networkName := "random_333_network"
	addNetworkCommand.Execute(&rpc, &networkId, &networkName)
}

func TestF_Should_Fail_To_Add_Network_With_Existing_Name(t *testing.T) {
	mock.SetUp(t)

	m := networkService.Init()
	s := storageService.Init()
	addNetworkCommand := network.AddNetwork{
		Locator: implementations.Locator{
			NetworkService: m,
			Storage:        s,
		},
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but function did not panic")

		}
	}()

	rpc := "HTTPS://127.0.0.1:2245"
	networkId := 5777
	networkName := "random_123_network"

	addNetworkCommand.Execute(&rpc, &networkId, &networkName)
}
