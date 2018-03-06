package infrastructure

import (
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

// TransactionsBuilderFactory represents a concrete TransactionsBuilderFactory implementation
type TransactionsBuilderFactory struct {
}

// CreateTransactionsBuilderFactory creates a new TransactionsBuilderFactory instance
func CreateTransactionsBuilderFactory() stored_signed_transactions.TransactionsBuilderFactory {
	out := TransactionsBuilderFactory{}
	return &out
}

// Create creates a new TransactionsBuilder instance
func (fac *TransactionsBuilderFactory) Create() stored_signed_transactions.TransactionsBuilder {
	out := createTransactionsBuilder()
	return out
}
