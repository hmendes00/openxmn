package infrastructure

import (
	"errors"
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	signed "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
	concrete_signed "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/infrastructure"
	uuid "github.com/satori/go.uuid"
)

type aggregatedTransactionsBuilder struct {
	htBuilderFactory hashtrees.HashTreeBuilderFactory
	id               *uuid.UUID
	trs              []signed.Transaction
	atomicTrs        []signed.AtomicTransaction
	createdOn        *time.Time
}

func createTransactionsBuilder(htBuilderFactory hashtrees.HashTreeBuilderFactory) aggregated.TransactionsBuilder {
	out := aggregatedTransactionsBuilder{
		htBuilderFactory: htBuilderFactory,
		id:               nil,
		trs:              nil,
		atomicTrs:        nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the TransactionsBuilder instance
func (build *aggregatedTransactionsBuilder) Create() aggregated.TransactionsBuilder {
	build.id = nil
	build.trs = []signed.Transaction{}
	build.atomicTrs = []signed.AtomicTransaction{}
	build.createdOn = nil
	return build
}

// WithID adds an  ID to the TransactionsBuilder
func (build *aggregatedTransactionsBuilder) WithID(id *uuid.UUID) aggregated.TransactionsBuilder {
	build.id = id
	return build
}

// WithTransactions adds signed Transactions to the TransactionsBuilder
func (build *aggregatedTransactionsBuilder) WithTransactions(trs []signed.Transaction) aggregated.TransactionsBuilder {
	build.trs = trs
	return build
}

// WithAtomicTransactions adds signed AtomicTransactions to the TransactionsBuilder
func (build *aggregatedTransactionsBuilder) WithAtomicTransactions(trs []signed.AtomicTransaction) aggregated.TransactionsBuilder {
	build.atomicTrs = trs
	return build
}

// CreatedOn adds a creation time to the TransactionsBuilder
func (build *aggregatedTransactionsBuilder) CreatedOn(ts time.Time) aggregated.TransactionsBuilder {
	build.createdOn = &ts
	return build
}

// Now builds an Transactions instance
func (build *aggregatedTransactionsBuilder) Now() (aggregated.Transactions, error) {

	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a transactions instance")
	}

	if build.createdOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a transactions instance")
	}

	if build.atomicTrs == nil && build.trs == nil {
		return nil, errors.New("there is no transactions or atomic transactions, therefore the aggregated transactions cannot be built")
	}

	//declare the hashtree block:
	htBlocks := [][]byte{}

	//transactions:
	trs := []*concrete_signed.Transaction{}
	if build.trs != nil {
		for _, oneTrs := range build.trs {
			//add the ID in the block:
			htBlocks = append(htBlocks, oneTrs.GetID().Bytes())

			//append the atomic trs:
			trs = append(trs, oneTrs.(*concrete_signed.Transaction))
		}
	}

	//atomic transactions:
	atomicTrs := []*concrete_signed.AtomicTransaction{}
	if build.atomicTrs != nil {
		for _, oneAtomicTrs := range build.atomicTrs {
			//add the ID in the block:
			htBlocks = append(htBlocks, oneAtomicTrs.GetID().Bytes())

			//append the atomic trs:
			atomicTrs = append(atomicTrs, oneAtomicTrs.(*concrete_signed.AtomicTransaction))
		}
	}

	//build the hashtree:
	ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(htBlocks).Now()
	if htErr != nil {
		return nil, htErr
	}

	if len(trs) > 0 && len(atomicTrs) > 0 {
		out := createTransactions(build.id, ht.(*concrete_hashtrees.HashTree), trs, atomicTrs, *build.createdOn)
		return out, nil
	}

	if len(atomicTrs) > 0 {
		out := createTransactionsWithAtomicTrs(build.id, ht.(*concrete_hashtrees.HashTree), atomicTrs, *build.createdOn)
		return out, nil
	}

	out := createTransactionsWithTrs(build.id, ht.(*concrete_hashtrees.HashTree), trs, *build.createdOn)
	return out, nil
}
