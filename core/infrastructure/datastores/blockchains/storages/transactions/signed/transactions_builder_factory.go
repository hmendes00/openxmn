package signed

import (
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions/signed"
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
