package infrastructure

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
)

// SignedTransactions represents the concrete signed transactions
type SignedTransactions struct {
	Met *concrete_metadata.MetaData `json:"metadata"`
	Trs *Transactions               `json:"transactions"`
	Sig *concrete_users.Signature   `json:"signature"`
}

func createSignedTransactions(met *concrete_metadata.MetaData, trs *Transactions, sig *concrete_users.Signature) aggregated.SignedTransactions {
	out := SignedTransactions{
		Met: met,
		Trs: trs,
		Sig: sig,
	}

	return &out
}

// GetMetaData returns the MetaData
func (trs *SignedTransactions) GetMetaData() metadata.MetaData {
	return trs.Met
}

// GetTransactions returns the Transactions instance
func (trs *SignedTransactions) GetTransactions() aggregated.Transactions {
	return trs.Trs
}

// GetSignature returns the user Signature
func (trs *SignedTransactions) GetSignature() users.Signature {
	return trs.Sig
}
