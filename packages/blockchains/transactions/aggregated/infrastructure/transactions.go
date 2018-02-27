package infrastructure

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	signed "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
	concrete_signed "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// Transactions represents the concrete aggregated transactions
type Transactions struct {
	ID        *uuid.UUID                           `json:"id"`
	HT        *concrete_hashtrees.HashTree         `json:"hashtree"`
	Trs       []*concrete_signed.Transaction       `json:"transactions"`
	AtomicTrs []*concrete_signed.AtomicTransaction `json:"atomic_transactions"`
	CrOn      time.Time                            `json:"created_on"`
}

func createTransactions(id *uuid.UUID, ht *concrete_hashtrees.HashTree, trs []*concrete_signed.Transaction, atomicTrs []*concrete_signed.AtomicTransaction, createdOn time.Time) aggregated.Transactions {
	out := Transactions{
		ID:        id,
		HT:        ht,
		Trs:       trs,
		AtomicTrs: atomicTrs,
		CrOn:      createdOn,
	}

	return &out
}

func createTransactionsWithTrs(id *uuid.UUID, ht *concrete_hashtrees.HashTree, trs []*concrete_signed.Transaction, createdOn time.Time) aggregated.Transactions {
	out := Transactions{
		ID:        id,
		HT:        ht,
		Trs:       trs,
		AtomicTrs: nil,
		CrOn:      createdOn,
	}

	return &out
}

func createTransactionsWithAtomicTrs(id *uuid.UUID, ht *concrete_hashtrees.HashTree, atomicTrs []*concrete_signed.AtomicTransaction, createdOn time.Time) aggregated.Transactions {
	out := Transactions{
		ID:        id,
		HT:        ht,
		Trs:       nil,
		AtomicTrs: atomicTrs,
		CrOn:      createdOn,
	}

	return &out
}

// GetID returns the ID
func (trs *Transactions) GetID() *uuid.UUID {
	return trs.ID
}

// GetHashTree returns the HashTree
func (trs *Transactions) GetHashTree() hashtrees.HashTree {
	return trs.HT
}

// HasTrs returns true if there is transaction, false otherwise
func (trs *Transactions) HasTrs() bool {
	return trs.Trs != nil
}

// GetTrs returns the signed transactions
func (trs *Transactions) GetTrs() []signed.Transaction {
	out := []signed.Transaction{}
	for _, oneTrs := range trs.Trs {
		out = append(out, oneTrs)
	}

	return out
}

// HasAtomicTrs returns true if there is an atomic transaction, false otherwise
func (trs *Transactions) HasAtomicTrs() bool {
	return trs.AtomicTrs != nil
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
