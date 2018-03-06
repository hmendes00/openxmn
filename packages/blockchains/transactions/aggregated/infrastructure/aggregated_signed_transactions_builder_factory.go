package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
)

// AggregatedSignedTransactionsBuilderFactory represents a concrete AggregatedSignedTransactionsBuilderFactory implementation
type AggregatedSignedTransactionsBuilderFactory struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
}

// CreateAggregatedSignedTransactionsBuilderFactory creates a new AggregatedSignedTransactionsBuilderFactory instance
func CreateAggregatedSignedTransactionsBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) aggregated.AggregatedSignedTransactionsBuilderFactory {
	out := AggregatedSignedTransactionsBuilderFactory{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
	}

	return &out
}

// Create creates a new AggregatedSignedTransactionsBuilder instance
func (fac *AggregatedSignedTransactionsBuilderFactory) Create() aggregated.AggregatedSignedTransactionsBuilder {
	out := createAggregatedSignedTransactionsBuilder(fac.htBuilderFactory, fac.metaDataBuilderFactory)
	return out
}
