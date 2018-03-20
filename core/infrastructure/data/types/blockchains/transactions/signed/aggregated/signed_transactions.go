package aggregated

import (
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	aggregated "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions/signed/aggregated"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/users"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/metadata"
	concrete_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/users"
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
