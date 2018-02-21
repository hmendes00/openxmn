package infrastructure

import (
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
)

// TransactionBuilderFactory represents a TransactionBuilderFactory implementation
type TransactionBuilderFactory struct {
}

// CreateTransactionBuilderFactory creates a new TransactionBuilderFactory instance
func CreateTransactionBuilderFactory() stored_transactions.TransactionBuilderFactory {
	out := TransactionBuilderFactory{}
	return &out
}

// Create creates a new TransactionBuilder instance
func (fac *TransactionBuilderFactory) Create() stored_transactions.TransactionBuilder {
	out := createTransactionBuilder()
	return out
}
