package infrastructure

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

type transactions struct {
	metaData  stored_files.File
	ht        stored_files.File
	trs       []stored_signed_transactions.Transaction
	atomicTrs []stored_signed_transactions.AtomicTransaction
}

func createTransactions(metaData stored_files.File, ht stored_files.File, trs []stored_signed_transactions.Transaction, atomicTrs []stored_signed_transactions.AtomicTransaction) stored_aggregated_transactions.Transactions {
	out := transactions{
		metaData:  metaData,
		ht:        ht,
		trs:       trs,
		atomicTrs: atomicTrs,
	}

	return &out
}

// GetMetaData returns the metadata file
func (trs *transactions) GetMetaData() stored_files.File {
	return trs.metaData
}

// GetHashTree returns the hashtree file
func (trs *transactions) GetHashTree() stored_files.File {
	return trs.ht
}

// HasTrs returns true if there is stored transactions, false otherwise
func (trs *transactions) HasTrs() bool {
	return trs.trs != nil
}

// GetTrs returns the stored transactions
func (trs *transactions) GetTrs() []stored_signed_transactions.Transaction {
	return trs.trs
}

// HasAtomicTrs returns true if there is stored atomic transactions, false otherwise
func (trs *transactions) HasAtomicTrs() bool {
	return trs.atomicTrs != nil
}

// GetAtomicTrs returns the stored atomic transactions
func (trs *transactions) GetAtomicTrs() []stored_signed_transactions.AtomicTransaction {
	return trs.atomicTrs
}
