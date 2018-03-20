package signed

import (
	stored_signed_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions/signed"
)

// AtomicTransactionsBuilderFactory represents a concrete AtomicTransactionsBuilderFactory implementation
type AtomicTransactionsBuilderFactory struct {
}

// CreateAtomicTransactionsBuilderFactory creates a new AtomicTransactionsBuilderFactory instance
func CreateAtomicTransactionsBuilderFactory() stored_signed_transactions.AtomicTransactionsBuilderFactory {
	out := AtomicTransactionsBuilderFactory{}
	return &out
}

// Create creates a new AtomicTransactions instance
func (fac *AtomicTransactionsBuilderFactory) Create() stored_signed_transactions.AtomicTransactionsBuilder {
	out := createAtomicTransactionsBuilder()
	return out
}
