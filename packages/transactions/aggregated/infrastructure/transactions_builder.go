package infrastructure

import (
	"errors"

	aggregated "github.com/XMNBlockchain/core/packages/transactions/aggregated/domain"
	signed "github.com/XMNBlockchain/core/packages/transactions/signed/domain"
	concrete_signed "github.com/XMNBlockchain/core/packages/transactions/signed/infrastructure"
)

type aggregatedTransactionsBuilder struct {
	trs       []signed.Transaction
	atomicTrs []signed.AtomicTransaction
}

func createTransactionsBuilder() aggregated.TransactionsBuilder {
	out := aggregatedTransactionsBuilder{
		trs:       []signed.Transaction{},
		atomicTrs: []signed.AtomicTransaction{},
	}

	return &out
}

// Create initializes the TransactionsBuilder instance
func (build *aggregatedTransactionsBuilder) Create() aggregated.TransactionsBuilder {
	build.trs = []signed.Transaction{}
	build.atomicTrs = []signed.AtomicTransaction{}
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

// Now builds an Transactions instance
func (build *aggregatedTransactionsBuilder) Now() (aggregated.Transactions, error) {

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

	out := createTransactions(trs, atomicTrs)
	return out, nil
}
