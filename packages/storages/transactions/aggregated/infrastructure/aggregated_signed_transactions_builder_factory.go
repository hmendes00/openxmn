package infrastructure

import (
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
)

// AggregatedSignedTransactionsBuilderFactory represents a concrete AggregatedSignedTransactionsBuilderFactory implementation
type AggregatedSignedTransactionsBuilderFactory struct {
}

// CreateAggregatedSignedTransactionsBuilderFactory creates a new AggregatedSignedTransactionsBuilderFactory instance
func CreateAggregatedSignedTransactionsBuilderFactory() stored_aggregated_transactions.AggregatedSignedTransactionsBuilderFactory {
	out := AggregatedSignedTransactionsBuilderFactory{}
	return &out
}

// Create creates a new AggregatedSignedTransactionsBuilder instance
func (fac *AggregatedSignedTransactionsBuilderFactory) Create() stored_aggregated_transactions.AggregatedSignedTransactionsBuilder {
	out := createAggregatedSignedTransactionsBuilder()
	return out
}
