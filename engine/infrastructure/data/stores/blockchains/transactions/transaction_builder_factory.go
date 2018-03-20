package transactions

import (
	stored_transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/blockchains/transactions"
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
