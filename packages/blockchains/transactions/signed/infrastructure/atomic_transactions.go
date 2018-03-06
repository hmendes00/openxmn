package infrastructure

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	signed_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
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
