package signed

import (
	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	signed_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed"
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
