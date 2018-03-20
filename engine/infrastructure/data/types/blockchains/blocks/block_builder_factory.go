package blocks

import (
	blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks"
	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
)

// BlockBuilderFactory represents a concrete BlockBuilderFactory implementation
type BlockBuilderFactory struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
}

// CreateBlockBuilderFactory creates a new BlockBuilderFactory instance
func CreateBlockBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) blocks.BlockBuilderFactory {
	out := BlockBuilderFactory{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
	}

	return &out
}

// Create creates a new BlockBuilder instance
func (fac *BlockBuilderFactory) Create() blocks.BlockBuilder {
	out := createBlockBuilder(fac.htBuilderFactory, fac.metaDataBuilderFactory)
	return out
}
