package aggregated

import (
	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	aggregated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed/aggregated"
)

// TransactionsBuilderFactory represents a concrete TransactionsBuilder instance
type TransactionsBuilderFactory struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.BuilderFactory
}

// CreateTransactionsBuilderFactory creates a new TransactionsBuilderFactory instance
func CreateTransactionsBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.BuilderFactory) aggregated.TransactionsBuilderFactory {
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
