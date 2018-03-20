package blocks

import (
	blocks "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/blocks"
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/hashtrees"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/metadata"
)

// SignedBlockBuilderFactory represents a concrete SignedBlockBuilderFactory implementation
type SignedBlockBuilderFactory struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
}

// CreateSignedBlockBuilderFactory creates a SignedBlockBuilderFactory instance
func CreateSignedBlockBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) blocks.SignedBlockBuilderFactory {
	out := SignedBlockBuilderFactory{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
	}
	return &out
}

// Create creates a new SignedBlockBuilder instance
func (fac *SignedBlockBuilderFactory) Create() blocks.SignedBlockBuilder {
	out := createSignedBlockBuilder(fac.htBuilderFactory, fac.metaDataBuilderFactory)
	return out
}
