package core

import "bytes"

// TXInput represents transaction input
type TXInput struct {
	// prev transaction id
	Txid []byte
	// index of an output in prev transaction
	Vout int
	// Signature of input transaction
	Signature []byte
	// Public key of wallet
	PubKey []byte
}

func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}
