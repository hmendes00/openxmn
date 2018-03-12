package transactions

import (
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions"
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
