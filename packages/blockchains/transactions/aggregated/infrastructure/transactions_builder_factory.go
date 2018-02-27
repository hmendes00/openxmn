package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
)

// TransactionsBuilderFactory represents a concrete TransactionsBuilder instance
type TransactionsBuilderFactory struct {
	htBuilderFactory hashtrees.HashTreeBuilderFactory
}

// CreateTransactionsBuilderFactory creates a new TransactionsBuilderFactory instance
func CreateTransactionsBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory) aggregated.TransactionsBuilderFactory {
	out := TransactionsBuilderFactory{
		htBuilderFactory: htBuilderFactory,
	}
	return &out
}

// Create creates an AgregatedTransactionsBuilder instance
func (fac *TransactionsBuilderFactory) Create() aggregated.TransactionsBuilder {
	out := createTransactionsBuilder(fac.htBuilderFactory)
	return out
}
