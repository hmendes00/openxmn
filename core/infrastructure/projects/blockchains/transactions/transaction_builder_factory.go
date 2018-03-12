package transactions

import (
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/hashtrees"
	met "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	trs "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions"
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
