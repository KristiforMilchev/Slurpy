package services_test

import (
	"os"
	"testing"

	"Slurpy/implementations"
	"Slurpy/interfaces"
)

var encoder interfaces.EncoderService

func Test_Should_Parse_Valid_ABI(t *testing.T) {
	encoder = &implementations.EncoderService{}

	file, err := os.ReadFile("../data/example_abi.json")

	if err != nil {
		t.Log("Local example json doesn't exist, please check test data in setup!")
		t.FailNow()
	}

	abi, err := encoder.ParseAbi(&file)

	if abi == nil || err != nil {
		t.Log(err)
		t.Log("Failed to parse ABI, expected ABI but got nil")
		t.Fail()
	}
}

func Test_Should_Fail_To_Parse_Empty_ABI(t *testing.T) {
	encoder = &implementations.EncoderService{}
	var empty []byte
	_, err := encoder.ParseAbi(&empty)
	if err == nil {
		t.Log("Nil should be rejected when parsing ABI")
		t.FailNow()
	}
}

func Test_Parse_Bytecode_From_Valid_Hex(t *testing.T) {
	encoder = &implementations.EncoderService{}
	bytecode := "0x6080604052348015600f57600080fd5b5060405160c838038060c8833981016040819052602a91604e565b600080546001600160a01b0319166001600160a01b0392909216919091179055607c565b600060208284031215605f57600080fd5b81516001600160a01b0381168114607557600080fd5b9392505050565b603f8060896000396000f3fe6080604052600080fdfea2646970667358221220164ced733355373d04d2603f54a7b92e4ba5abdfc29a3db4485038f2f79773a064736f6c63430008130033"

	_, err := encoder.ParseByteCode(&bytecode)

	if err != nil {
		t.Log(err)
		t.Log("Failed to parse valid bytecode in hex format!")
		t.Fail()
	}
}

func Test_Parse_Bytecode_From_Valid_Hex_Without_0x(t *testing.T) {
	encoder = &implementations.EncoderService{}
	bytecode := "6080604052348015600f57600080fd5b5060405160c838038060c8833981016040819052602a91604e565b600080546001600160a01b0319166001600160a01b0392909216919091179055607c565b600060208284031215605f57600080fd5b81516001600160a01b0381168114607557600080fd5b9392505050565b603f8060896000396000f3fe6080604052600080fdfea2646970667358221220164ced733355373d04d2603f54a7b92e4ba5abdfc29a3db4485038f2f79773a064736f6c63430008130033"

	_, err := encoder.ParseByteCode(&bytecode)

	if err != nil {
		t.Log(err)
		t.Log("Failed to parse valid bytecode in hex format!")
		t.Fail()
	}
}

func Test_Error_If_Invalid_Hex_Is_Provided(t *testing.T) {
	encoder = &implementations.EncoderService{}
	invalidHex := []string{
		"0xGHIJKL",   // Invalid characters
		"0x123Z56",   // Invalid characters
		"HelloWorld", // Only letters
		"0x1",        // Odd length
		//"0x12345",    // Odd length
		"",          // Empty string
		" ",         // Whitespace
		"0xHello",   // Valid prefix, invalid characters
		"0x123_456", // Invalid character
	}

	for _, hex := range invalidHex {
		_, err := encoder.ParseByteCode(&hex)

		if err == nil {
			t.Log(hex, "Should be invalid")
			t.Log("Parsed illegal hex, should have returned an exception")
			t.FailNow()
		}
	}
}
