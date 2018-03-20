package transactions

import (
	stored_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions"
)

// BuilderFactory represents a concrete BuilderFactory implementation
type BuilderFactory struct {
}

// CreateBuilderFactory creates a new BuilderFactory instance
func CreateBuilderFactory() stored_transactions.BuilderFactory {
	out := BuilderFactory{}
	return &out
}

// Create creates a new TransactionsBuilder instance
func (fac *BuilderFactory) Create() stored_transactions.Builder {
	out := createBuilder()
	return out
}
