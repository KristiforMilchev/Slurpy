package interfaces

import "github.com/ethereum/go-ethereum/accounts/abi"

type EncoderService interface {
	ParseAbi(abi *[]byte) (*abi.ABI, error)
	ParseByteCode(bytecode *string) ([]byte, error)
}
