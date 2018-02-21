package infrastructure

import (
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
)

// TransactionsBuilderFactory represents a concrete TransactionsBuilderFactory implementation
type TransactionsBuilderFactory struct {
}

// CreateTransactionsBuilderFactory creates a new TransactionsBuilderFactory instance
func CreateTransactionsBuilderFactory() stored_aggregated_transactions.TransactionsBuilderFactory {
	out := TransactionsBuilderFactory{}
	return &out
}

// Create creates a new TransactionsBuilder instance
func (fac *TransactionsBuilderFactory) Create() stored_aggregated_transactions.TransactionsBuilder {
	out := createTransactionsBuilder()
	return out
}
