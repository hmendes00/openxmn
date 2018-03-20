package signed

import (
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/hashtrees"
	met "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/metadata"
	signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/transactions/signed"
)

// AtomicTransactionBuilderFactory represents a concrete AtomicTransactionBuilder factory
type AtomicTransactionBuilderFactory struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory met.MetaDataBuilderFactory
}

// CreateAtomicTransactionBuilderFactory creates a new AtomicTransactionBuilderFactory instance
func CreateAtomicTransactionBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory met.MetaDataBuilderFactory) signed_transactions.AtomicTransactionBuilderFactory {
	out := AtomicTransactionBuilderFactory{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
	}
	return &out
}

// Create creates a new AtomicTransactionBuilder instance
func (fac *AtomicTransactionBuilderFactory) Create() signed_transactions.AtomicTransactionBuilder {
	out := createAtomicTransactionBuilder(fac.htBuilderFactory, fac.metaDataBuilderFactory)
	return out
}
