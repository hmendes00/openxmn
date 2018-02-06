package infrastructure

import (
	signed_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
)

// TransactionBuilderFactory represents a concrete TransactionBuilder factory
type TransactionBuilderFactory struct {
}

// CreateTransactionBuilderFactory creates a new TransactionBuilderFactory instance
func CreateTransactionBuilderFactory() signed_transactions.TransactionBuilderFactory {
	out := TransactionBuilderFactory{}
	return &out
}

// Create creates a new TransactionBuilder instance
func (fac *TransactionBuilderFactory) Create() signed_transactions.TransactionBuilder {
	out := createTransactionBuilder()
	return out
}
