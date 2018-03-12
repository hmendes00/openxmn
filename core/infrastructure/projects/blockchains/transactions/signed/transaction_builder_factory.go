package signed

import (
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/hashtrees"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions/signed"
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
