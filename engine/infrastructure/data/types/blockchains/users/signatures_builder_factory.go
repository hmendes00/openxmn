package users

import (
	hashtrees "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/hashtrees"
	metadata "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/metadata"
	user "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/users"
)

// SignaturesBuilderFactory represents a concrete SignaturesBuilderFactory implementation
type SignaturesBuilderFactory struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
}

// CreateSignaturesBuilderFactory creates a new SignaturesBuilderFactory instance
func CreateSignaturesBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) user.SignaturesBuilderFactory {
	out := SignaturesBuilderFactory{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
	}

	return &out
}

// Create builds a new SignaturesBuilder instance
func (fac *SignaturesBuilderFactory) Create() user.SignaturesBuilder {
	out := createSignaturesBuilder(fac.htBuilderFactory, fac.metaDataBuilderFactory)
	return out
}
