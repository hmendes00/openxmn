package infrastructure

import (
	"errors"

	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

type transactionsBuilder struct {
	metaData  stored_files.File
	ht        stored_files.File
	trs       []stored_signed_transactions.Transaction
	atomicTrs []stored_signed_transactions.AtomicTransaction
}

func createTransactionsBuilder() stored_aggregated_transactions.TransactionsBuilder {
	out := transactionsBuilder{
		metaData:  nil,
		ht:        nil,
		trs:       nil,
		atomicTrs: nil,
	}

	return &out
}

// Create initializes the TransactionsBuilder instance
func (build *transactionsBuilder) Create() stored_aggregated_transactions.TransactionsBuilder {
	build.metaData = nil
	build.ht = nil
	build.trs = nil
	build.atomicTrs = nil
	return build
}

// WithMetaData adds metadata file to the TransactionsBuilder instance
func (build *transactionsBuilder) WithMetaData(met stored_files.File) stored_aggregated_transactions.TransactionsBuilder {
	build.metaData = met
	return build
}

// WithHashTree adds an hashtree file to the TransactionsBuilder instance
func (build *transactionsBuilder) WithHashTree(ht stored_files.File) stored_aggregated_transactions.TransactionsBuilder {
	build.ht = ht
	return build
}

// WithTrs adds stored transactions to the TransactionsBuilder instance
func (build *transactionsBuilder) WithTrs(trs []stored_signed_transactions.Transaction) stored_aggregated_transactions.TransactionsBuilder {
	build.trs = trs
	return build
}

// WithAtomicTrs adds stored atomic transactions to the TransactionsBuilder instance
func (build *transactionsBuilder) WithAtomicTrs(atomicTrs []stored_signed_transactions.AtomicTransaction) stored_aggregated_transactions.TransactionsBuilder {
	build.atomicTrs = atomicTrs
	return build
}

// Now builds a Transactions instance
func (build *transactionsBuilder) Now() (stored_aggregated_transactions.Transactions, error) {
	if build.metaData == nil {
		return nil, errors.New("the metadata file is mandatory in order to build a Transactions instance")
	}

	if build.ht == nil {
		return nil, errors.New("the hashtree file is mandatory in order to build a Transactions instance")
	}

	if build.trs == nil {
		return nil, errors.New("the transactions file is mandatory in order to build a Transactions instance")
	}

	if build.atomicTrs == nil {
		return nil, errors.New("the atomic transactions file is mandatory in order to build a Transactions instance")
	}

	out := createTransactions(build.metaData, build.ht, build.trs, build.atomicTrs)
	return out, nil
}