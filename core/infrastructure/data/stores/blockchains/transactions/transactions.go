package transactions

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/files"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/transactions"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/files"
)

// Transactions represents a concrete stored transactions implemetation
type Transactions struct {
	Met *concrete_stored_files.File `json:"metadata"`
	Trs []*Transaction              `json:"transactions"`
}

func createTransactions(met *concrete_stored_files.File, trs []*Transaction) stored_transactions.Transactions {
	out := Transactions{
		Met: met,
		Trs: trs,
	}

	return &out
}

// GetMetaData returns the metadata
func (trs *Transactions) GetMetaData() stored_files.File {
	return trs.Met
}

// GetTransactions returns the transactions
func (trs *Transactions) GetTransactions() []stored_transactions.Transaction {
	out := []stored_transactions.Transaction{}
	for _, oneTrs := range trs.Trs {
		out = append(out, oneTrs)
	}
	return out
}
