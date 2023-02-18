package core

import (
	"bytes"
	"mini-chain/core/utils"
)

// TXOutput represents transaction output
type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

func (out *TXOutput) Lock(address []byte) {
	decodedAddress := utils.Base58Decode(address)
	// decodedAddress = <version><pubKeyHash><checkSum>
	// version is 1 byte
	// checkSum is 4 bytes
	out.PubKeyHash = decodedAddress[1 : len(decodedAddress)-4]
}

func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// NewTXOutput create a new TXOutput
// value is transaction value in native currency of this chain
// address is base58 encoded wallet address of receiver of the transaction
func NewTXOutput(value int, address string) *TXOutput {
	txo := &TXOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}
