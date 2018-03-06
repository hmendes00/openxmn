package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
)

// TransactionsBuilderFactory represents a concrete Transactions implementation
type TransactionsBuilderFactory struct {
	htBuilderFactory  hashtrees.HashTreeBuilderFactory
	metBuilderFactory metadata.MetaDataBuilderFactory
	trs               []trs.Transaction
}

// CreateTransactionsBuilderFactory creates a new TransactionsBuilderFactory instance
func CreateTransactionsBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metBuilderFactory metadata.MetaDataBuilderFactory) transactions.TransactionsBuilderFactory {
	out := TransactionsBuilderFactory{
		htBuilderFactory:  htBuilderFactory,
		metBuilderFactory: metBuilderFactory,
	}

	return &out
}

// Create initializes the TransactionsBuilder instance
func (fac *TransactionsBuilderFactory) Create() transactions.TransactionsBuilder {
	out := createTransactionsBuilder(fac.htBuilderFactory, fac.metBuilderFactory)
	return out
}
