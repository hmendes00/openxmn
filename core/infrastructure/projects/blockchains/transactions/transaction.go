package transactions

import (
	met "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	trs "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions"
	concrete_met "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/metadata"
)

// Transaction represents the concrete transaction
type Transaction struct {
	Met *concrete_met.MetaData `json:"metadata"`
	JS  []byte                 `json:"json"`
}

func createTransaction(met *concrete_met.MetaData, js []byte) trs.Transaction {
	out := Transaction{
		Met: met,
		JS:  js,
	}

	return &out
}

// GetMetaData returns the MetaData of the transaction
func (trs *Transaction) GetMetaData() met.MetaData {
	return trs.Met
}

// GetJSON returns the json data of the transaction
func (trs *Transaction) GetJSON() []byte {
	return trs.JS
}
