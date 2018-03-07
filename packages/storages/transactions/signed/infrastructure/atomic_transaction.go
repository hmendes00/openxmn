package infrastructure

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
)

type atomicTransaction struct {
	metaData stored_files.File
	sig      stored_users.Signature
	trs      stored_transactions.Transactions
}

func createAtomicTransaction(metaData stored_files.File, sig stored_users.Signature, trs stored_transactions.Transactions) stored_signed_transactions.AtomicTransaction {
	out := atomicTransaction{
		metaData: metaData,
		sig:      sig,
		trs:      trs,
	}

	return &out
}

// GetMetaData returns the metadata file
func (trs *atomicTransaction) GetMetaData() stored_files.File {
	return trs.metaData
}

// GetSignature returns the signature file
func (trs *atomicTransaction) GetSignature() stored_users.Signature {
	return trs.sig
}

// GetTransactions returns the stored transactions
func (trs *atomicTransaction) GetTransactions() stored_transactions.Transactions {
	return trs.trs
}
