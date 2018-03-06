package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	signed_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
)

// TransactionsBuilderFactory represents a concrete TransactionsBuilderFactory implementation
type TransactionsBuilderFactory struct {
	htBuilderFactory  hashtrees.HashTreeBuilderFactory
	metBuilderFactory metadata.MetaDataBuilderFactory
}

// CreateTransactionsBuilderFactory creates a new TransactionsBuilderFactory instance
func CreateTransactionsBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metBuilderFactory metadata.MetaDataBuilderFactory) signed_transactions.TransactionsBuilderFactory {
	out := TransactionsBuilderFactory{
		htBuilderFactory:  htBuilderFactory,
		metBuilderFactory: metBuilderFactory,
	}

	return &out
}

// Create creates a new Transactions instance
func (fac *TransactionsBuilderFactory) Create() signed_transactions.TransactionsBuilder {
	out := createTransactionsBuilder(fac.htBuilderFactory, fac.metBuilderFactory)
	return out
}
