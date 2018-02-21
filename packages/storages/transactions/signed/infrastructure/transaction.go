package infrastructure

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
)

type transaction struct {
	metaData stored_files.File
	sig      stored_files.File
	trs      stored_transactions.Transaction
}

func createTransaction(metaData stored_files.File, sig stored_files.File, trs stored_transactions.Transaction) stored_signed_transactions.Transaction {
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

// GetSignature returns the signature file
func (trs *transaction) GetSignature() stored_files.File {
	return trs.sig
}

// GetTransaction returns the stored transaction
func (trs *transaction) GetTransaction() stored_transactions.Transaction {
	return trs.trs
}
