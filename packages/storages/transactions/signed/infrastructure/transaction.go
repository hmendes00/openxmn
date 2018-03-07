package infrastructure

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
)

type transaction struct {
	metaData stored_files.File
	sig      stored_users.Signature
	trs      stored_transactions.Transaction
}

func createTransaction(metaData stored_files.File, sig stored_users.Signature, trs stored_transactions.Transaction) stored_signed_transactions.Transaction {
	out := transaction{
		metaData: metaData,
		sig:      sig,
		trs:      trs,
	}

	return &out
}

// GetMetaData returns the metadata file
func (trs *transaction) GetMetaData() stored_files.File {
	return trs.metaData
}

// GetSignature returns the signature
func (trs *transaction) GetSignature() stored_users.Signature {
	return trs.sig
}

// GetTransaction returns the stored transaction
func (trs *transaction) GetTransaction() stored_transactions.Transaction {
	return trs.trs
}
