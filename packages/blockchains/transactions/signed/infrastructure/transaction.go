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

// Transaction represents the concrete signed transaction
type Transaction struct {
	Met *concrete_metadata.MetaData        `json:"metadata"`
	Trs *concrete_transactions.Transaction `json:"transaction"`
	Sig *concrete_users.Signature          `json:"signature"`
}

func createTransaction(met *concrete_metadata.MetaData, trs *concrete_transactions.Transaction, sig *concrete_users.Signature) signed_transactions.Transaction {
	out := Transaction{
		Met: met,
		Trs: trs,
		Sig: sig,
	}

	return &out
}

// GetMetaData returns the MetaData
func (trs *Transaction) GetMetaData() met.MetaData {
	return trs.Met
}

// GetTransaction returns the Transaction
func (trs *Transaction) GetTransaction() trs.Transaction {
	return trs.Trs
}

// GetSignature returns the user signature
func (trs *Transaction) GetSignature() users.Signature {
	return trs.Sig
}
