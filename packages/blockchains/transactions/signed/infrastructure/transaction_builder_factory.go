package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	signed_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
)

// TransactionBuilderFactory represents a concrete TransactionBuilder factory
type TransactionBuilderFactory struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
}

// CreateTransactionBuilderFactory creates a new TransactionBuilderFactory instance
func CreateTransactionBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) signed_transactions.TransactionBuilderFactory {
	out := TransactionBuilderFactory{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
	}
	return &out
}

// Create creates a new TransactionBuilder instance
func (fac *TransactionBuilderFactory) Create() signed_transactions.TransactionBuilder {
	out := createTransactionBuilder(fac.htBuilderFactory, fac.metaDataBuilderFactory)
	return out
}
