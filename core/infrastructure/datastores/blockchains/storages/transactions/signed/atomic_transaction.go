package signed

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions"
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions/signed"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/users"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/files"
	conrete_stored_transactions "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/transactions"
	conrete_stored_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/users"
)

// AtomicTransaction represents a concrete stored AtomicTransaction implementation
type AtomicTransaction struct {
	MetaData *concrete_stored_files.File               `json:"metadata"`
	Sig      *conrete_stored_users.Signature           `json:"signature"`
	Trs      *conrete_stored_transactions.Transactions `json:"transactions"`
}

func createAtomicTransaction(metaData *concrete_stored_files.File, sig *conrete_stored_users.Signature, trs *conrete_stored_transactions.Transactions) stored_signed_transactions.AtomicTransaction {
	out := AtomicTransaction{
		MetaData: metaData,
		Sig:      sig,
		Trs:      trs,
	}

	return &out
}

// GetMetaData returns the metadata file
func (trs *AtomicTransaction) GetMetaData() stored_files.File {
	return trs.MetaData
}

// GetSignature returns the signature file
func (trs *AtomicTransaction) GetSignature() stored_users.Signature {
	return trs.Sig
}

// GetTransactions returns the stored transactions
func (trs *AtomicTransaction) GetTransactions() stored_transactions.Transactions {
	return trs.Trs
}
