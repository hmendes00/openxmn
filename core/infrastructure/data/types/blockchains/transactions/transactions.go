package transactions

import (
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions"
	trs "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/metadata"
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
