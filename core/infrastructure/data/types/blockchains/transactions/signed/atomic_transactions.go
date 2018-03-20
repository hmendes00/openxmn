package signed

import (
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions/signed"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/metadata"
)

// AtomicTransactions represents a concrete AtomicTransactions implemetation
type AtomicTransactions struct {
	Met *concrete_metadata.MetaData `json:"metadata"`
	Trs []*AtomicTransaction        `json:"atomic_transactions"`
}

func createAtomicTransactions(met *concrete_metadata.MetaData, trs []*AtomicTransaction) signed_transactions.AtomicTransactions {
	out := AtomicTransactions{
		Met: met,
		Trs: trs,
	}

	return &out
}

// GetMetaData returns the MetaData
func (trs *AtomicTransactions) GetMetaData() metadata.MetaData {
	return trs.Met
}

// GetTransactions returns the []AtomicTransaction
func (trs *AtomicTransactions) GetTransactions() []signed_transactions.AtomicTransaction {
	out := []signed_transactions.AtomicTransaction{}
	for _, oneTrs := range trs.Trs {
		out = append(out, oneTrs)
	}

	return out
}
