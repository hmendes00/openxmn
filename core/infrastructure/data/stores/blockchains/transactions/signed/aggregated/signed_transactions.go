package aggregated

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/files"
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/transactions/signed/aggregated"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/users"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/files"
	concrete_stored_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/blockchains/users"
)

// SignedTransactions represents a concrete stored aggregated SignedTransactions implementation
type SignedTransactions struct {
	MetaData *concrete_stored_files.File      `json:"metadata"`
	Sig      *concrete_stored_users.Signature `json:"signature"`
	Trs      *Transactions                    `json:"aggregated_transactions"`
}

func createSignedTransactions(metaData *concrete_stored_files.File, sig *concrete_stored_users.Signature, trs *Transactions) stored_aggregated_transactions.SignedTransactions {
	out := SignedTransactions{
		MetaData: metaData,
		Sig:      sig,
		Trs:      trs,
	}

	return &out
}

// GetMetaData returns the metadata file
func (trs *SignedTransactions) GetMetaData() stored_files.File {
	return trs.MetaData
}

// GetSignature returns the signature file
func (trs *SignedTransactions) GetSignature() stored_users.Signature {
	return trs.Sig
}

// GetTransactions returns the stored transactions
func (trs *SignedTransactions) GetTransactions() stored_aggregated_transactions.Transactions {
	return trs.Trs
}
