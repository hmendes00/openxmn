package infrastructure

import (
	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	hashtree "github.com/XMNBlockchain/core/packages/hashtrees/domain"
)

// BlockBuilderFactory represents a concrete BlockBuilderFactory implementation
type BlockBuilderFactory struct {
	htBuilderFactory hashtree.HashTreeBuilderFactory
}

// CreateBlockBuilderFactory creates a BlockBuilderFactory instance
func CreateBlockBuilderFactory(htBuilderFactory hashtree.HashTreeBuilderFactory) blocks.BlockBuilderFactory {
	out := BlockBuilderFactory{
		htBuilderFactory: htBuilderFactory,
	}

	return &out
}

// Create creates a new BlockBuilder instance
func (fac *BlockBuilderFactory) Create() blocks.BlockBuilder {
	out := createBlockBuilder(fac.htBuilderFactory)
	return out
}
