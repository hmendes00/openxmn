package aggregated

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions/signed/aggregated"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/users"
)

type signedTransactions struct {
	metaData stored_files.File
	sig      stored_users.Signature
	trs      stored_aggregated_transactions.Transactions
}

func createSignedTransactions(metaData stored_files.File, sig stored_users.Signature, trs stored_aggregated_transactions.Transactions) stored_aggregated_transactions.SignedTransactions {
	out := signedTransactions{
		metaData: metaData,
		sig:      sig,
		trs:      trs,
	}

	return &out
}

// GetMetaData returns the metadata file
func (trs *signedTransactions) GetMetaData() stored_files.File {
	return trs.metaData
}

// GetSignature returns the signature file
func (trs *signedTransactions) GetSignature() stored_users.Signature {
	return trs.sig
}

// GetTransactions returns the stored transactions
func (trs *signedTransactions) GetTransactions() stored_aggregated_transactions.Transactions {
	return trs.trs
}
