package infrastructure

import (
	stored_chained_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/chained/domain"
)

// BlockBuilderFactory represents a concrete BlockBuilderFactory implementation
type BlockBuilderFactory struct {
}

// CreateBlockBuilderFactory creates a new BlockBuilderFactory instance
func CreateBlockBuilderFactory() stored_chained_blocks.BlockBuilderFactory {
	out := BlockBuilderFactory{}
	return &out
}

// Create creates a new BlockBuilder instance
func (fac *BlockBuilderFactory) Create() stored_chained_blocks.BlockBuilder {
	out := createBlockBuilder()
	return out
}
