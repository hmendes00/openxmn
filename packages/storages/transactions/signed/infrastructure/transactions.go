package infrastructure

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

type transactions struct {
	met stored_files.File
	trs []stored_signed_transactions.Transaction
}

func createTransactions(met stored_files.File, trs []stored_signed_transactions.Transaction) stored_signed_transactions.Transactions {
	out := transactions{
		met: met,
		trs: trs,
	}

	return &out
}

// GetMetaData returns the MetaData
func (trs *transactions) GetMetaData() stored_files.File {
	return trs.met
}

// GetTransactions returns the []Transaction
func (trs *transactions) GetTransactions() []stored_signed_transactions.Transaction {
	return trs.trs
}