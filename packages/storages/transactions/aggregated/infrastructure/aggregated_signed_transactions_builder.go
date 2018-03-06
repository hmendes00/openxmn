package infrastructure

import (
	"errors"

	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
)

type aggregatedSignedTransactionsBuilder struct {
	met stored_files.File
	trs []stored_aggregated_transactions.SignedTransactions
}

func createAggregatedSignedTransactionsBuilder() stored_aggregated_transactions.AggregatedSignedTransactionsBuilder {
	out := aggregatedSignedTransactionsBuilder{
		met: nil,
		trs: nil,
	}

	return &out
}

// Create initializes the AggregatedSignedTransactionsBuilder
func (build *aggregatedSignedTransactionsBuilder) Create() stored_aggregated_transactions.AggregatedSignedTransactionsBuilder {
	build.met = nil
	build.trs = nil
	return build
}

// WithMetaData adds MetaData to the AggregatedSignedTransactionsBuilder instance
func (build *aggregatedSignedTransactionsBuilder) WithMetaData(met stored_files.File) stored_aggregated_transactions.AggregatedSignedTransactionsBuilder {
	build.met = met
	return build
}

// WithTransactions adds SignedTransactions to the AggregatedSignedTransactionsBuilder instance
func (build *aggregatedSignedTransactionsBuilder) WithTransactions(trs []stored_aggregated_transactions.SignedTransactions) stored_aggregated_transactions.AggregatedSignedTransactionsBuilder {
	build.trs = trs
	return build
}

// Now builds a new AggregatedSignedTransactions instance
func (build *aggregatedSignedTransactionsBuilder) Now() (stored_aggregated_transactions.AggregatedSignedTransactions, error) {
	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build an AggregatedSignedTransactions instance")
	}

	if build.trs == nil {
		return nil, errors.New("the SignedTransactions are mandatory in order to build an AggregatedSignedTransactions instance")
	}

	out := createAggregatedSignedTransactions(build.met, build.trs)
	return out, nil
}
