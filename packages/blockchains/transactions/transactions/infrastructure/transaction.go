package infrastructure

import (
	met "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	concrete_met "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
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
