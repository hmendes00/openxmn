package signed

import (
	met "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	trs "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	signed_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_transactions "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/transactions"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
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
