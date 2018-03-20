package signed

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/files"
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/blockchains/transactions/signed"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/stores/files"
)

// AtomicTransactions represents a concrete stored AtomicTransactions implementation
type AtomicTransactions struct {
	Met *concrete_stored_files.File `json:"metadata"`
	Trs []*AtomicTransaction        `json:"atomic_transactions"`
}

func createAtomicTransactions(met *concrete_stored_files.File, trs []*AtomicTransaction) stored_signed_transactions.AtomicTransactions {
	out := AtomicTransactions{
		Met: met,
		Trs: trs,
	}

	return &out
}

// GetMetaData returns the MetaData
func (trs *AtomicTransactions) GetMetaData() stored_files.File {
	return trs.Met
}

// GetTransactions returns the []AtomicTransaction
func (trs *AtomicTransactions) GetTransactions() []stored_signed_transactions.AtomicTransaction {
	out := []stored_signed_transactions.AtomicTransaction{}
	for _, oneTrs := range trs.Trs {
		out = append(out, oneTrs)
	}
	return out
}
