package signed

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions"
	stored_signed_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions/signed"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/users"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
	concrete_stored_transactions "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/blockchains/transactions"
	concrete_stored_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/users"
)

// Transaction represents a concrete stored signed transaction implementation
type Transaction struct {
	MetaData *concrete_stored_files.File               `json:"metadata"`
	Sig      *concrete_stored_users.Signature          `json:"signature"`
	Trs      *concrete_stored_transactions.Transaction `json:"transaction"`
}

func createTransaction(metaData *concrete_stored_files.File, sig *concrete_stored_users.Signature, trs *concrete_stored_transactions.Transaction) stored_signed_transactions.Transaction {
	out := Transaction{
		MetaData: metaData,
		Sig:      sig,
		Trs:      trs,
	}

	return &out
}

// GetMetaData returns the metadata file
func (trs *Transaction) GetMetaData() stored_files.File {
	return trs.MetaData
}

// GetSignature returns the signature
func (trs *Transaction) GetSignature() stored_users.Signature {
	return trs.Sig
}

// GetTransaction returns the stored transaction
func (trs *Transaction) GetTransaction() stored_transactions.Transaction {
	return trs.Trs
}
