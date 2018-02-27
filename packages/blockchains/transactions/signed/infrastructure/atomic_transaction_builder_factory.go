package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	signed_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
)

// AtomicTransactionBuilderFactory represents a concrete AtomicTransactionBuilder factory
type AtomicTransactionBuilderFactory struct {
	htBuilderFactory hashtrees.HashTreeBuilderFactory
}

// CreateAtomicTransactionBuilderFactory creates a new AtomicTransactionBuilderFactory instance
func CreateAtomicTransactionBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory) signed_transactions.AtomicTransactionBuilderFactory {
	out := AtomicTransactionBuilderFactory{
		htBuilderFactory: htBuilderFactory,
	}
	return &out
}

// Create creates a new AtomicTransactionBuilder instance
func (fac *AtomicTransactionBuilderFactory) Create() signed_transactions.AtomicTransactionBuilder {
	out := createAtomicTransactionBuilder(fac.htBuilderFactory)
	return out
}
