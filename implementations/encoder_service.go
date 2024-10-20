package implementations

import (
	"bytes"
	"encoding/hex"
	"errors"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type EncoderService struct {
}

func (e *EncoderService) ParseAbi(abiData *[]byte) (*abi.ABI, error) {
	if abiData == nil || len(*abiData) == 0 {
		return nil, errors.New("nil encoded data is not allowed")
	}
	byteReader := bytes.NewReader(*abiData)
	parsedABI, err := abi.JSON(byteReader)
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
		return nil, err
	}

	return &parsedABI, nil
}

func (e *EncoderService) ParseByteCode(bytecodeHex *string) ([]byte, error) {
	if *bytecodeHex == "" {
		return nil, errors.New("empty Strings are not allowed")
	}

	cleanedHex := strings.TrimPrefix(*bytecodeHex, "0x")
	bytecode, err := hex.DecodeString(string(cleanedHex))
	if err != nil {
		return nil, err
	}

	return bytecode, nil
}
