package infrastructure

import (
	met "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	signed_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
	trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	concrete_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/infrastructure"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
)

// AtomicTransaction represents the concrete signed atomic transaction
type AtomicTransaction struct {
	Met *concrete_metadata.MetaData         `json:"metadata"`
	Trs *concrete_transactions.Transactions `json:"transactions"`
	Sig *concrete_users.Signature           `json:"signature"`
}

func createAtomicTransaction(meta *concrete_metadata.MetaData, trs *concrete_transactions.Transactions, sig *concrete_users.Signature) signed_transactions.AtomicTransaction {
	out := AtomicTransaction{
		Met: meta,
		Trs: trs,
		Sig: sig,
	}

	return &out
}

// GetMetaData returns the transaction MetaData
func (atomic *AtomicTransaction) GetMetaData() met.MetaData {
	return atomic.Met
}

// GetTransactions returns the transactions
func (atomic *AtomicTransaction) GetTransactions() trs.Transactions {
	return atomic.Trs
}

// GetSignature returns the user signature, if any
func (atomic *AtomicTransaction) GetSignature() users.Signature {
	return atomic.Sig
}
