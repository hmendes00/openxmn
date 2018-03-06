package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
)

// TransactionsBuilderFactory represents a concrete TransactionsBuilder instance
type TransactionsBuilderFactory struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
}

// CreateTransactionsBuilderFactory creates a new TransactionsBuilderFactory instance
func CreateTransactionsBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) aggregated.TransactionsBuilderFactory {
	out := TransactionsBuilderFactory{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
	}
	return &out
}

// Create creates an AgregatedTransactionsBuilder instance
func (fac *TransactionsBuilderFactory) Create() aggregated.TransactionsBuilder {
	out := createTransactionsBuilder(fac.htBuilderFactory, fac.metaDataBuilderFactory)
	return out
}
