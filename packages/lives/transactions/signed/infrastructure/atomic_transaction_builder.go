package infrastructure

import (
	"errors"
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	signed_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	trs "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	concrete_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

type atomicTransactionBuilder struct {
	htBuilderFactory hashtrees.HashTreeBuilderFactory
	id               *uuid.UUID
	trs              []trs.Transaction
	sig              users.Signature
	createdOn        *time.Time
}

func createAtomicTransactionBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory) signed_transactions.AtomicTransactionBuilder {
	out := atomicTransactionBuilder{
		htBuilderFactory: htBuilderFactory,
		id:               nil,
		trs:              nil,
		sig:              nil,
		createdOn:        nil,
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

// CreatedOn adds a creation time to the AtomicTransactionBuilder
func (build *atomicTransactionBuilder) CreatedOn(ts time.Time) signed_transactions.AtomicTransactionBuilder {
	build.createdOn = &ts
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

	if build.createdOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build an AtomicTransaction instance")
	}

	htBlocks := [][]byte{}
	trs := []*concrete_transactions.Transaction{}
	for _, oneTrs := range build.trs {
		//add the ID in the block:
		htBlocks = append(htBlocks, oneTrs.GetID().Bytes())

		//add the trs in the list:
		trs = append(trs, oneTrs.(*concrete_transactions.Transaction))
	}

	//build the hashtree:
	ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(htBlocks).Now()
	if htErr != nil {
		return nil, htErr
	}

	out := createAtomicTransaction(build.id, ht.(*concrete_hashtrees.HashTree), trs, build.sig.(*concrete_users.Signature), *build.createdOn)
	return out, nil
}
