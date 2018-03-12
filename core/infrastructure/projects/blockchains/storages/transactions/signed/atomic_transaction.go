package signed

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions"
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions/signed"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/users"
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
