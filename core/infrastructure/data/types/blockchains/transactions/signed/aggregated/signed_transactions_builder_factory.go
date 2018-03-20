package aggregated

import (
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/hashtrees"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	aggregated "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions/signed/aggregated"
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
