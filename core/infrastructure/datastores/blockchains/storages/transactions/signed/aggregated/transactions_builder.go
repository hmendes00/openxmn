package aggregated

import (
	"errors"

	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions/signed"
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions/signed/aggregated"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/files"
	concrete_stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/transactions/signed"
)

type transactionsBuilder struct {
	metaData  stored_files.File
	trs       stored_signed_transactions.Transactions
	atomicTrs stored_signed_transactions.AtomicTransactions
}

func createTransactionsBuilder() stored_aggregated_transactions.TransactionsBuilder {
	out := transactionsBuilder{
		metaData:  nil,
		trs:       nil,
		atomicTrs: nil,
	}

	return &out
}

// Create initializes the TransactionsBuilder instance
func (build *transactionsBuilder) Create() stored_aggregated_transactions.TransactionsBuilder {
	build.metaData = nil
	build.trs = nil
	build.atomicTrs = nil
	return build
}

// WithMetaData adds metadata file to the TransactionsBuilder instance
func (build *transactionsBuilder) WithMetaData(met stored_files.File) stored_aggregated_transactions.TransactionsBuilder {
	build.metaData = met
	return build
}

// WithTrs adds stored transactions to the TransactionsBuilder instance
func (build *transactionsBuilder) WithTransactions(trs stored_signed_transactions.Transactions) stored_aggregated_transactions.TransactionsBuilder {
	build.trs = trs
	return build
}

// WithAtomicTrs adds stored atomic transactions to the TransactionsBuilder instance
func (build *transactionsBuilder) WithAtomicTransactions(atomicTrs stored_signed_transactions.AtomicTransactions) stored_aggregated_transactions.TransactionsBuilder {
	build.atomicTrs = atomicTrs
	return build
}

// Now builds a Transactions instance
func (build *transactionsBuilder) Now() (stored_aggregated_transactions.Transactions, error) {
	if build.metaData == nil {
		return nil, errors.New("the metadata file is mandatory in order to build a Transactions instance")
	}

	if build.trs == nil {
		return nil, errors.New("the transactions file is mandatory in order to build a Transactions instance")
	}

	if build.atomicTrs == nil {
		return nil, errors.New("the atomic transactions file is mandatory in order to build a Transactions instance")
	}

	out := createTransactions(build.metaData.(*concrete_stored_files.File), build.trs.(*concrete_stored_signed_transactions.Transactions), build.atomicTrs.(*concrete_stored_signed_transactions.AtomicTransactions))
	return out, nil
}
