package implementations_test

import (
	"testing"

	"github.com/ovechkin-dm/mockio/mock"

	"slurpy/implementations"
	"slurpy/interfaces"
)

var config interfaces.Configuration

func TestF_Load_Should_Return_True(t *testing.T) {
	mock.SetUp(t)
	filePath := "../test_data/settings.json"
	config = &implementations.Configuration{
		File: &filePath,
	}

	hasLoaded := config.Load()

	if !hasLoaded {
		t.Log("Configuration should not fail to load with existing configuration file")
		t.FailNow()
	}
}

func TestF_Load_Should_Panic_On_Malformed_Configurtaion_File(t *testing.T) {
	mock.SetUp(t)
	filePath := "../test_data/malformed_setting.json"
	config = &implementations.Configuration{
		File: &filePath,
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but function did not panic")
		}
	}()

	config.Load()
}

func TestF_Load_Should_Panic_On_Missing_Configuration_File(t *testing.T) {
	mock.SetUp(t)
	filePath := "000x000"
	config = &implementations.Configuration{
		File: &filePath,
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but function did not panic")
		}
	}()

	config.Load()
}

func TestF_Should_Return_Existing_Key(t *testing.T) {
	mock.SetUp(t)
	filePath := "../test_data/settings.json"
	config = &implementations.Configuration{
		File: &filePath,
	}

	config.Load()

	dbName := config.GetKey("DbName")
	if dbName == nil {
		t.Log("Failed to retrive value from settings.json with existing key")
		t.FailNow()
	}
}

func TestF_Should_Return_Nil_For_Missing_Key(t *testing.T) {
	mock.SetUp(t)
	filePath := "../test_data/settings.json"
	config = &implementations.Configuration{
		File: &filePath,
	}

	config.Load()

	dbName := config.GetKey("this_is_missing")
	if dbName != nil {
		t.Log("Failed to retrive value from settings.json with existing key")
		t.FailNow()
	}
}
