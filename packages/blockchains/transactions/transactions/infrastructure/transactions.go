package infrastructure

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
)

// Transactions represents a concrete Transactions implementation
type Transactions struct {
	Met *concrete_metadata.MetaData `json:"metadata"`
	Trs []*Transaction              `json:"transactions"`
}

func createTransactions(met *concrete_metadata.MetaData, trs []*Transaction) transactions.Transactions {
	out := Transactions{
		Met: met,
		Trs: trs,
	}

	return &out
}

// GetMetaData returns the MetaData
func (atomic *Transactions) GetMetaData() metadata.MetaData {
	return atomic.Met
}

// GetTransactions returns the transactions
func (atomic *Transactions) GetTransactions() []trs.Transaction {
	out := []trs.Transaction{}
	for _, oneTrs := range atomic.Trs {
		out = append(out, oneTrs)
	}

	return out
}
