package transactions

import (
	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	trs "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
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
