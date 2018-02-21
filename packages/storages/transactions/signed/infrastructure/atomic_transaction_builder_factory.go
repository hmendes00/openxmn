package infrastructure

import (
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

// AtomicTransactionBuilderFactory represents a concrete AtomicTransactionBuilderFactory implementation
type AtomicTransactionBuilderFactory struct {
}

// CreateAtomicTransactionBuilderFactory creates a new AtomicTransactionBuilderFactory instance
func CreateAtomicTransactionBuilderFactory() stored_signed_transactions.AtomicTransactionBuilderFactory {
	out := AtomicTransactionBuilderFactory{}
	return &out
}

// Create creates a new AtomicTransactionBuilder instance
func (fac *AtomicTransactionBuilderFactory) Create() stored_signed_transactions.AtomicTransactionBuilder {
	out := createAtomicTransactionBuilder()
	return out
}
