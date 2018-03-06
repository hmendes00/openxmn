package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	signed_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
)

// AtomicTransactionsBuilderFactory represents a concrete AtomicTransactionsBuilderFactory implementation
type AtomicTransactionsBuilderFactory struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
}

// CreateAtomicTransactionsBuilderFactory creates a new AtomicTransactionsBuilderFactory instance
func CreateAtomicTransactionsBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) signed_transactions.AtomicTransactionsBuilderFactory {
	out := AtomicTransactionsBuilderFactory{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
	}

	return &out
}

// Create creates a new AtomicTransactionsBuilder instance
func (fac *AtomicTransactionsBuilderFactory) Create() signed_transactions.AtomicTransactionsBuilder {
	out := createAtomicTransactionsBuilder(fac.htBuilderFactory, fac.metaDataBuilderFactory)
	return out
}
