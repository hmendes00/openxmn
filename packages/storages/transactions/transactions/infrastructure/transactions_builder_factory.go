package infrastructure

import (
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
)

// TransactionsBuilderFactory represents a concrete TransactionsBuilderFactory implementation
type TransactionsBuilderFactory struct {
}

// CreateTransactionsBuilderFactory creates a new TransactionsBuilderFactory instance
func CreateTransactionsBuilderFactory() stored_transactions.TransactionsBuilderFactory {
	out := TransactionsBuilderFactory{}
	return &out
}

// Create creates a new TransactionsBuilder instance
func (fac *TransactionsBuilderFactory) Create() stored_transactions.TransactionsBuilder {
	out := createTransactionsBuilder()
	return out
}
