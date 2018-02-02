package infrastructure

import (
	signed_transactions "github.com/XMNBlockchain/core/packages/transactions/signed/domain"
)

// AtomicTransactionBuilderFactory represents a concrete AtomicTransactionBuilder factory
type AtomicTransactionBuilderFactory struct {
}

// CreateAtomicTransactionBuilderFactory creates a new AtomicTransactionBuilderFactory instance
func CreateAtomicTransactionBuilderFactory() signed_transactions.AtomicTransactionBuilderFactory {
	out := AtomicTransactionBuilderFactory{}
	return &out
}

// Create creates a new AtomicTransactionBuilder instance
func (fac *AtomicTransactionBuilderFactory) Create() signed_transactions.AtomicTransactionBuilder {
	out := createAtomicTransactionBuilder()
	return out
}
