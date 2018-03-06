package infrastructure

import (
	"errors"

	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

type transactionsBuilder struct {
	met stored_files.File
	trs []stored_signed_transactions.Transaction
}

func createTransactionsBuilder() stored_signed_transactions.TransactionsBuilder {
	out := transactionsBuilder{
		met: nil,
		trs: nil,
	}

	return &out
}

// Create initializes the TransactionsBuilder instance
func (build *transactionsBuilder) Create() stored_signed_transactions.TransactionsBuilder {
	build.met = nil
	build.trs = nil
	return build
}

// WithMetaData adds MetaData to the TransactionsBuilder instance
func (build *transactionsBuilder) WithMetaData(met stored_files.File) stored_signed_transactions.TransactionsBuilder {
	build.met = met
	return build
}

// WithTransactions adds []Transaction to the TransactionsBuilder instance
func (build *transactionsBuilder) WithTransactions(trs []stored_signed_transactions.Transaction) stored_signed_transactions.TransactionsBuilder {
	build.trs = trs
	return build
}

// Now builds a new Transactions instance
func (build *transactionsBuilder) Now() (stored_signed_transactions.Transactions, error) {
	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a Transactions instance")
	}

	if build.trs == nil {
		return nil, errors.New("the []Transaction is mandatory in order to build a Transactions instance")
	}

	out := createTransactions(build.met, build.trs)
	return out, nil
}
