package transactions

import (
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions"
)

// BuilderFactory represents a concrete BuilderFactory implementation
type BuilderFactory struct {
}

// CreateBuilderFactory creates a new BuilderFactory instance
func CreateBuilderFactory() stored_transactions.TransactionsBuilderFactory {
	out := BuilderFactory{}
	return &out
}

// Create creates a new TransactionsBuilder instance
func (fac *BuilderFactory) Create() stored_transactions.TransactionsBuilder {
	out := createBuilder()
	return out
}
