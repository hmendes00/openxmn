package aggregated

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions/signed"
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions/signed/aggregated"
)

type transactions struct {
	metaData  stored_files.File
	trs       stored_signed_transactions.Transactions
	atomicTrs stored_signed_transactions.AtomicTransactions
}

func createTransactions(metaData stored_files.File, trs stored_signed_transactions.Transactions, atomicTrs stored_signed_transactions.AtomicTransactions) stored_aggregated_transactions.Transactions {
	out := transactions{
		metaData:  metaData,
		trs:       trs,
		atomicTrs: atomicTrs,
	}

	return &out
}

// GetMetaData returns the metadata file
func (trs *transactions) GetMetaData() stored_files.File {
	return trs.metaData
}

// HasTrs returns true if there is stored transactions, false otherwise
func (trs *transactions) HasTransactions() bool {
	return trs.trs != nil
}

// GetTrs returns the stored transactions
func (trs *transactions) GetTransactions() stored_signed_transactions.Transactions {
	return trs.trs
}

// HasAtomicTrs returns true if there is stored atomic transactions, false otherwise
func (trs *transactions) HasAtomicTransactions() bool {
	return trs.atomicTrs != nil
}

// GetAtomicTrs returns the stored atomic transactions
func (trs *transactions) GetAtomicTransactions() stored_signed_transactions.AtomicTransactions {
	return trs.atomicTrs
}
