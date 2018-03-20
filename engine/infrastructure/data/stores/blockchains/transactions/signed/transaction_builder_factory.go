package signed

import (
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/blockchains/transactions/signed"
)

// TransactionBuilderFactory represents a concrete TransactionBuilderFactory implementation
type TransactionBuilderFactory struct {
}

// CreateTransactionBuilderFactory creates a new TransactionBuilderFactory instance
func CreateTransactionBuilderFactory() stored_signed_transactions.TransactionBuilderFactory {
	out := TransactionBuilderFactory{}
	return &out
}

// Create creates a new TransactionBuilder instance
func (fac *TransactionBuilderFactory) Create() stored_signed_transactions.TransactionBuilder {
	out := createTransactionBuilder()
	return out
}
