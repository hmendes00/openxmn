package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
)

// SignedTransactionsBuilderFactory represents the concrete SignedTransactionsBuilder factory
type SignedTransactionsBuilderFactory struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
}

// CreateSignedTransactionsBuilderFactory creates a new SignedTransactionsBuilderFactory instance
func CreateSignedTransactionsBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) aggregated.SignedTransactionsBuilderFactory {
	out := SignedTransactionsBuilderFactory{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
	}
	return &out
}

// Create creates a new SignedTransactionsBuilder instance
func (fac *SignedTransactionsBuilderFactory) Create() aggregated.SignedTransactionsBuilder {
	out := createSignedTransactionsBuilder(fac.htBuilderFactory, fac.metaDataBuilderFactory)
	return out
}
