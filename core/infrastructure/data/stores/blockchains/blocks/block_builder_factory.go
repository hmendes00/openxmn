package blocks

import (
	stored_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/blocks"
)

// BlockBuilderFactory represents a concrete BlockBuilderFactory implementation
type BlockBuilderFactory struct {
}

// CreateBlockBuilderFactory creates a new BlockBuilderFactory instance
func CreateBlockBuilderFactory() stored_blocks.BlockBuilderFactory {
	out := BlockBuilderFactory{}
	return &out
}

// Create creates a new BlockBuilder instance
func (fac *BlockBuilderFactory) Create() stored_blocks.BlockBuilder {
	out := createBlockBuilder()
	return out
}
