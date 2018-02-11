package infrastructure

import (
	"errors"
	"time"

	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	signed "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	concrete_signed "github.com/XMNBlockchain/core/packages/lives/transactions/signed/infrastructure"
	uuid "github.com/satori/go.uuid"
)

type aggregatedTransactionsBuilder struct {
	id        *uuid.UUID
	trs       []signed.Transaction
	atomicTrs []signed.AtomicTransaction
	createdOn *time.Time
}

func createTransactionsBuilder() aggregated.TransactionsBuilder {
	out := aggregatedTransactionsBuilder{
		id:        nil,
		trs:       []signed.Transaction{},
		atomicTrs: []signed.AtomicTransaction{},
		createdOn: nil,
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

	if len(build.atomicTrs) <= 0 && len(build.trs) <= 0 {
		return nil, errors.New("there is no transactions or atomic transactions, therefore the aggregated transactions cannot be built")
	}

	trs := []*concrete_signed.Transaction{}
	for _, oneTrs := range build.trs {
		trs = append(trs, oneTrs.(*concrete_signed.Transaction))
	}

	atomicTrs := []*concrete_signed.AtomicTransaction{}
	for _, oneAtomicTrs := range build.atomicTrs {
		atomicTrs = append(atomicTrs, oneAtomicTrs.(*concrete_signed.AtomicTransaction))
	}

	out := createTransactions(build.id, trs, atomicTrs, *build.createdOn)
	return out, nil
}
