package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
)

// BlockBuilderFactory represents a concrete BlockBuilderFactory implementation
type BlockBuilderFactory struct {
	htBuilderFactory hashtrees.HashTreeBuilderFactory
}

// CreateBlockBuilderFactory creates a BlockBuilderFactory instance
func CreateBlockBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory) blocks.BlockBuilderFactory {
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
