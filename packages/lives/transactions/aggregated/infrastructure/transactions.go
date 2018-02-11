package infrastructure

import (
	"time"

	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	signed "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	concrete_signed "github.com/XMNBlockchain/core/packages/lives/transactions/signed/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// Transactions represents the concrete aggregated transactions
type Transactions struct {
	ID        *uuid.UUID                           `json:"id"`
	Trs       []*concrete_signed.Transaction       `json:"transactions"`
	AtomicTrs []*concrete_signed.AtomicTransaction `json:"atomic_transactions"`
	CrOn      time.Time                            `json:"created_on"`
}

func createTransactions(id *uuid.UUID, trs []*concrete_signed.Transaction, atomicTrs []*concrete_signed.AtomicTransaction, createdOn time.Time) aggregated.Transactions {
	out := Transactions{
		ID:        id,
		Trs:       trs,
		AtomicTrs: atomicTrs,
		CrOn:      createdOn,
	}

	return &out
}

// GetID returns the ID
func (trs *Transactions) GetID() *uuid.UUID {
	return trs.ID
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

// CreatedOn returns the creation time
func (trs *Transactions) CreatedOn() time.Time {
	return trs.CrOn
}
