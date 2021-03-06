package aggregated

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	signed "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed"
	aggregated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed/aggregated"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_signed "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/transactions/signed"
)

// Transactions represents the concrete aggregated transactions
type Transactions struct {
	Met       *concrete_metadata.MetaData         `json:"metadata"`
	Trs       *concrete_signed.Transactions       `json:"transactions"`
	AtomicTrs *concrete_signed.AtomicTransactions `json:"atomic_transactions"`
}

func createTransactions(met *concrete_metadata.MetaData, trs *concrete_signed.Transactions, atomicTrs *concrete_signed.AtomicTransactions) aggregated.Transactions {
	out := Transactions{
		Met:       met,
		Trs:       trs,
		AtomicTrs: atomicTrs,
	}

	return &out
}

func createTransactionsWithTrs(met *concrete_metadata.MetaData, trs *concrete_signed.Transactions) aggregated.Transactions {
	out := Transactions{
		Met:       met,
		Trs:       trs,
		AtomicTrs: nil,
	}

	return &out
}

func createTransactionsWithAtomicTrs(met *concrete_metadata.MetaData, atomicTrs *concrete_signed.AtomicTransactions) aggregated.Transactions {
	out := Transactions{
		Met:       met,
		Trs:       nil,
		AtomicTrs: atomicTrs,
	}

	return &out
}

// GetMetaData returns the MetaData
func (trs *Transactions) GetMetaData() metadata.MetaData {
	return trs.Met
}

// HasTransactions returns true if there is transaction, false otherwise
func (trs *Transactions) HasTransactions() bool {
	return trs.Trs != nil
}

// GetTransactions returns the signed transactions
func (trs *Transactions) GetTransactions() signed.Transactions {
	return trs.Trs
}

// HasAtomicTransactions returns true if there is an atomic transaction, false otherwise
func (trs *Transactions) HasAtomicTransactions() bool {
	return trs.AtomicTrs != nil
}

// GetAtomicTransactions returns the signed atomic transactions
func (trs *Transactions) GetAtomicTransactions() signed.AtomicTransactions {
	return trs.AtomicTrs
}
