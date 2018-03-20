package signed

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_signed_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions/signed"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
)

// Transactions represents a concrete stored signed Transactions implementation
type Transactions struct {
	Met *concrete_stored_files.File `json:"metadata"`
	Trs []*Transaction              `json:"transactions"`
}

func createTransactions(met *concrete_stored_files.File, trs []*Transaction) stored_signed_transactions.Transactions {
	out := Transactions{
		Met: met,
		Trs: trs,
	}

	return &out
}

// GetMetaData returns the MetaData
func (trs *Transactions) GetMetaData() stored_files.File {
	return trs.Met
}

// GetTransactions returns the []Transaction
func (trs *Transactions) GetTransactions() []stored_signed_transactions.Transaction {
	out := []stored_signed_transactions.Transaction{}
	for _, oneTrs := range trs.Trs {
		out = append(out, oneTrs)
	}
	return out
}
