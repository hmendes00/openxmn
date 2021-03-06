package aggregated

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_signed_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions/signed"
	stored_aggregated_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions/signed/aggregated"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
	concrete_stored_signed_transactions "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/blockchains/transactions/signed"
)

// Transactions represents a concrete stored aggregated Transactions implementation
type Transactions struct {
	MetaData  *concrete_stored_files.File                             `json:"metadata"`
	Trs       *concrete_stored_signed_transactions.Transactions       `json:"signed_transactions"`
	AtomicTrs *concrete_stored_signed_transactions.AtomicTransactions `json:"signed_atomic_transactions"`
}

func createTransactions(metaData *concrete_stored_files.File, trs *concrete_stored_signed_transactions.Transactions, atomicTrs *concrete_stored_signed_transactions.AtomicTransactions) stored_aggregated_transactions.Transactions {
	out := Transactions{
		MetaData:  metaData,
		Trs:       trs,
		AtomicTrs: atomicTrs,
	}

	return &out
}

// GetMetaData returns the metadata file
func (trs *Transactions) GetMetaData() stored_files.File {
	return trs.MetaData
}

// HasTransactions returns true if there is stored transactions, false otherwise
func (trs *Transactions) HasTransactions() bool {
	return trs.Trs != nil
}

// GetTransactions returns the stored transactions
func (trs *Transactions) GetTransactions() stored_signed_transactions.Transactions {
	return trs.Trs
}

// HasAtomicTransactions returns true if there is stored atomic transactions, false otherwise
func (trs *Transactions) HasAtomicTransactions() bool {
	return trs.AtomicTrs != nil
}

// GetAtomicTransactions returns the stored atomic transactions
func (trs *Transactions) GetAtomicTransactions() stored_signed_transactions.AtomicTransactions {
	return trs.AtomicTrs
}
