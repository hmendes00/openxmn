package transactions

import (
	"errors"

	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions"
)

type transactionsBuilder struct {
	met stored_files.File
	trs []stored_transactions.Transaction
}

func createTransactionsBuilder() stored_transactions.TransactionsBuilder {
	out := transactionsBuilder{
		met: nil,
		trs: nil,
	}

	return &out
}

// Create initializes the Transactions bulder
func (build *transactionsBuilder) Create() stored_transactions.TransactionsBuilder {
	build.met = nil
	build.trs = nil
	return build
}

// WithMetaData adds MetaData to the Transactions builder
func (build *transactionsBuilder) WithMetaData(met stored_files.File) stored_transactions.TransactionsBuilder {
	build.met = met
	return build
}

// WithTransactions adds Transactions to the Transactions builder
func (build *transactionsBuilder) WithTransactions(trs []stored_transactions.Transaction) stored_transactions.TransactionsBuilder {
	build.trs = trs
	return build
}

// Now builds a new Transactions instance
func (build *transactionsBuilder) Now() (stored_transactions.Transactions, error) {
	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a stored Transactions instance")
	}

	if build.trs == nil {
		return nil, errors.New("the Transactions is mandatory in order to build a stored Transactions instance")
	}

	out := createTransactions(build.met, build.trs)
	return out, nil
}
