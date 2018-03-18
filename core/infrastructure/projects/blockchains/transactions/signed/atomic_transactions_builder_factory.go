package signed

import (
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/hashtrees"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions/signed"
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