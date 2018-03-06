package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	met "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
)

type transactionBuilderFactory struct {
	htBuilderFactory  hashtrees.HashTreeBuilderFactory
	metBuilderFactory met.MetaDataBuilderFactory
}

// CreateTransactionBuilderFactory creates a new TransactionBuilderFactory
func CreateTransactionBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metBuilderFactory met.MetaDataBuilderFactory) trs.TransactionBuilderFactory {
	out := transactionBuilderFactory{
		htBuilderFactory:  htBuilderFactory,
		metBuilderFactory: metBuilderFactory,
	}
	return &out
}

// Create creates a new TransactionBuilder instance
func (fac *transactionBuilderFactory) Create() trs.TransactionBuilder {
	out := createTransactionBuilder(fac.htBuilderFactory, fac.metBuilderFactory)
	return out
}
