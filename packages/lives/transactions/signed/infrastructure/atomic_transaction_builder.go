package infrastructure

import (
	"errors"
	"time"

	signed_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	concrete_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

type atomicTransactionBuilder struct {
	id  *uuid.UUID
	trs []trs.Transaction
	sig users.Signature
}

func createAtomicTransactionBuilder() signed_transactions.AtomicTransactionBuilder {
	out := atomicTransactionBuilder{
		id:  nil,
		trs: nil,
		sig: nil,
	}

	return &out
}

// Create initializes the AtomicTransactionBuilder instance
func (build *atomicTransactionBuilder) Create() signed_transactions.AtomicTransactionBuilder {
	build.id = nil
	build.trs = nil
	build.sig = nil
	return build
}

// WithID adds an ID to the AtomicTransactionBuilder
func (build *atomicTransactionBuilder) WithID(id *uuid.UUID) signed_transactions.AtomicTransactionBuilder {
	build.id = id
	return build
}

// WithTransactions adds transactions to the AtomicTransactionBuilder
func (build *atomicTransactionBuilder) WithTransactions(trs []trs.Transaction) signed_transactions.AtomicTransactionBuilder {
	build.trs = trs
	return build
}

// WithSignature adds a user signature to the AtomicTransactionBuilder
func (build *atomicTransactionBuilder) WithSignature(sig users.Signature) signed_transactions.AtomicTransactionBuilder {
	build.sig = sig
	return build
}

// Now builds a new AtomicTransaction instance
func (build *atomicTransactionBuilder) Now() (signed_transactions.AtomicTransaction, error) {

	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build an AtomicTransaction instance")
	}

	if build.trs == nil {
		return nil, errors.New("the []transaction are mandatory in order to build an AtomicTransaction instance")
	}

	if build.sig == nil {
		return nil, errors.New("the user signature is mandatory in order to build an AtomicTransaction instance")
	}

	if len(build.trs) <= 0 {
		return nil, errors.New("the []transaction cannot be empty")
	}

	isInit := false
	createdOn := time.Now()
	for _, oneTrs := range build.trs {

		if !isInit {
			createdOn = oneTrs.CreatedOn()
			isInit = true
			continue
		}

		oneCreatedOn := oneTrs.CreatedOn()
		if oneCreatedOn.Before(createdOn) {
			createdOn = oneCreatedOn
		}
	}

	trs := []*concrete_transactions.Transaction{}
	for _, oneTrs := range build.trs {
		trs = append(trs, oneTrs.(*concrete_transactions.Transaction))
	}

	out := createAtomicTransaction(build.id, trs, build.sig.(*concrete_users.Signature), createdOn)
	return out, nil
}
