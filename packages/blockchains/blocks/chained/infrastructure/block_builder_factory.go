package infrastructure

import (
	chained "github.com/XMNBlockchain/core/packages/blockchains/blocks/chained/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
)

// BlockBuilderFactory represents a concrete BlockBuilderFactory implementation
type BlockBuilderFactory struct {
	htBuilderFactory hashtrees.HashTreeBuilderFactory
}

// CreateBlockBuilderFactory creates a BlockBuilderFactory instance
func CreateBlockBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory) chained.BlockBuilderFactory {
	out := BlockBuilderFactory{
		htBuilderFactory: htBuilderFactory,
	}
	return &out
}

// Create creates a new BlockBuilder instance
func (fac *BlockBuilderFactory) Create() chained.BlockBuilder {
	out := createBlockBuilder(fac.htBuilderFactory)
	return out
}
