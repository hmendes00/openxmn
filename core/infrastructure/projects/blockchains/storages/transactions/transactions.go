package transactions

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions"
)

type transactions struct {
	met stored_files.File
	trs []stored_transactions.Transaction
}

func createTransactions(met stored_files.File, trs []stored_transactions.Transaction) stored_transactions.Transactions {
	out := transactions{
		met: met,
		trs: trs,
	}

	return &out
}

// GetMetaData returns the metadata
func (trs *transactions) GetMetaData() stored_files.File {
	return trs.met
}

// GetTransactions returns the transactions
func (trs *transactions) GetTransactions() []stored_transactions.Transaction {
	return trs.trs
}
