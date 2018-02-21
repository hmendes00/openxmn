package infrastructure

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
)

type atomicTransaction struct {
	metaData stored_files.File
	sig      stored_files.File
	ht       stored_files.File
	trs      []stored_transactions.Transaction
}

func createAtomicTransaction(metaData stored_files.File, sig stored_files.File, ht stored_files.File, trs []stored_transactions.Transaction) stored_signed_transactions.AtomicTransaction {
	out := atomicTransaction{
		metaData: metaData,
		sig:      sig,
		ht:       ht,
		trs:      trs,
	}

	return &out
}

// GetMetaData returns the metadata file
func (trs *atomicTransaction) GetMetaData() stored_files.File {
	return trs.metaData
}

// GetSignature returns the signature file
func (trs *atomicTransaction) GetSignature() stored_files.File {
	return trs.sig
}

// GetHashTree returns the hashtree file
func (trs *atomicTransaction) GetHashTree() stored_files.File {
	return trs.ht
}

// GetTransactions returns the stored transactions
func (trs *atomicTransaction) GetTransactions() []stored_transactions.Transaction {
	return trs.trs
}
