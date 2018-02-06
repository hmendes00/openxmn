package infrastructure

import (
	validated "github.com/XMNBlockchain/core/packages/lives/blocks/validated/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
)

// BlockBuilderFactory represents a concrete BlockBuilderFactory implementation
type BlockBuilderFactory struct {
	htBuilderFactory hashtrees.HashTreeBuilderFactory
}

// CreateBlockBuilderFactory creates a new BlockBuilderFactory instance
func CreateBlockBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory) validated.BlockBuilderFactory {
	out := BlockBuilderFactory{
		htBuilderFactory: htBuilderFactory,
	}

	return &out
}

// Create creates a new BlockBuilder instance
func (fac *BlockBuilderFactory) Create() validated.BlockBuilder {
	out := createBlockBuilder(fac.htBuilderFactory)
	return out
}
