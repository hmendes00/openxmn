package infrastructure

import (
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
)

// SignedTransactionsBuilderFactory represents a concrete SignedTransactionsBuilderFactory implementation
type SignedTransactionsBuilderFactory struct {
}

// CreateSignedTransactionsBuilderFactory creates a new SignedTransactionsBuilderFactory instance
func CreateSignedTransactionsBuilderFactory() stored_aggregated_transactions.SignedTransactionsBuilderFactory {
	out := SignedTransactionsBuilderFactory{}
	return &out
}

// Create creates a new SignedTransactionsBuilder instance
func (fac *SignedTransactionsBuilderFactory) Create() stored_aggregated_transactions.SignedTransactionsBuilder {
	out := createSignedTransactionsBuilder()
	return out
}
