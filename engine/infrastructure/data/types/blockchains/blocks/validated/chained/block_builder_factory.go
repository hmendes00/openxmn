package chained

import (
	chained "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated/chained"
	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
)

// BlockBuilderFactory represents a concrete BlockBuilderFactory implementation
type BlockBuilderFactory struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory chained.MetaDataBuilderFactory
}

// CreateBlockBuilderFactory creates a BlockBuilderFactory instance
func CreateBlockBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory chained.MetaDataBuilderFactory) chained.BlockBuilderFactory {
	out := BlockBuilderFactory{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
	}
	return &out
}

// Create creates a new BlockBuilder instance
func (fac *BlockBuilderFactory) Create() chained.BlockBuilder {
	out := createBlockBuilder(fac.htBuilderFactory, fac.metaDataBuilderFactory)
	return out
}
