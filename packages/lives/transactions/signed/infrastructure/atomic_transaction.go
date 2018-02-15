package infrastructure

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/infrastructure"
	signed_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	concrete_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// AtomicTransaction represents the concrete signed atomic transaction
type AtomicTransaction struct {
	ID   *uuid.UUID                           `json:"id"`
	HT   *concrete_hashtrees.HashTree         `json:"hashtree"`
	Trs  []*concrete_transactions.Transaction `json:"transactions"`
	Sig  *concrete_users.Signature            `json:"signature"`
	CrOn time.Time                            `json:"created_on"`
}

func createAtomicTransaction(id *uuid.UUID, ht *concrete_hashtrees.HashTree, trs []*concrete_transactions.Transaction, sig *concrete_users.Signature, createdOn time.Time) signed_transactions.AtomicTransaction {
	out := AtomicTransaction{
		ID:   id,
		HT:   ht,
		Trs:  trs,
		Sig:  sig,
		CrOn: createdOn,
	}

	return &out
}

// GetID returns the transaction ID
func (atomic *AtomicTransaction) GetID() *uuid.UUID {
	return atomic.ID
}

// GetHashTree returns the transactions HashTree
func (atomic *AtomicTransaction) GetHashTree() hashtrees.HashTree {
	return atomic.HT
}

// GetTrs returns the transactions
func (atomic *AtomicTransaction) GetTrs() []trs.Transaction {
	out := []trs.Transaction{}
	for _, oneTrs := range atomic.Trs {
		out = append(out, oneTrs)
	}

	return out
}

// GetSignature returns the user signature, if any
func (atomic *AtomicTransaction) GetSignature() users.Signature {
	return atomic.Sig
}

// CreatedOn returns the creation timestamp
func (atomic *AtomicTransaction) CreatedOn() time.Time {
	return atomic.CrOn
}
