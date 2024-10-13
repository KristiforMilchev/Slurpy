package implementations

import (
	"bytes"
	"encoding/hex"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type EncoderService struct {
}

func (e *EncoderService) ParseAbi(abiData *[]byte) (*abi.ABI, error) {
	byteReader := bytes.NewReader(*abiData)
	parsedABI, err := abi.JSON(byteReader)
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
		return nil, err
	}

	return &parsedABI, nil
}

func (e *EncoderService) ParseByteCode(bytecodeHex *string) ([]byte, error) {
	bytecode, err := hex.DecodeString(string(*bytecodeHex))
	if err != nil {
		log.Fatalf("Failed to decode bytecode: %v", err)
		return nil, err
	}

	return bytecode, nil
}
