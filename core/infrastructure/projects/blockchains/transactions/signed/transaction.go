package signed

import (
	met "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	trs "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions"
	signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions/signed"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/users"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/metadata"
	concrete_transactions "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/transactions"
	concrete_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/users"
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
