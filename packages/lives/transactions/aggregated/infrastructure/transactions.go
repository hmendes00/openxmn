package infrastructure

import (
	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	signed "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	concrete_signed "github.com/XMNBlockchain/core/packages/lives/transactions/signed/infrastructure"
)

// Transactions represents the concrete aggregated transactions
type Transactions struct {
	Trs       []*concrete_signed.Transaction       `json:"transactions"`
	AtomicTrs []*concrete_signed.AtomicTransaction `json:"atomic_transactions"`
}

func createTransactions(trs []*concrete_signed.Transaction, atomicTrs []*concrete_signed.AtomicTransaction) aggregated.Transactions {
	out := Transactions{
		Trs:       trs,
		AtomicTrs: atomicTrs,
	}

	return &out
}

// GetTrs returns the signed transactions
func (trs *Transactions) GetTrs() []signed.Transaction {
	out := []signed.Transaction{}
	for _, oneTrs := range trs.Trs {
		out = append(out, oneTrs)
	}

	return out
}

// GetAtomicTrs returns the signed atomic transactions
func (trs *Transactions) GetAtomicTrs() []signed.AtomicTransaction {
	out := []signed.AtomicTransaction{}
	for _, oneTrs := range trs.AtomicTrs {
		out = append(out, oneTrs)
	}

	return out
}
