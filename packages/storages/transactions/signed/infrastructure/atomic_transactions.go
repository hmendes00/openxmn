package infrastructure

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

type atomicTransactions struct {
	met stored_files.File
	trs []stored_signed_transactions.AtomicTransaction
}

func createAtomicTransactions(met stored_files.File, trs []stored_signed_transactions.AtomicTransaction) stored_signed_transactions.AtomicTransactions {
	out := atomicTransactions{
		met: met,
		trs: trs,
	}

	return &out
}

// GetMetaData returns the MetaData
func (trs *atomicTransactions) GetMetaData() stored_files.File {
	return trs.met
}

// GetTransactions returns the []AtomicTransaction
func (trs *atomicTransactions) GetTransactions() []stored_signed_transactions.AtomicTransaction {
	return trs.trs
}
