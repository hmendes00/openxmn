package infrastructure

import (
	validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/domain"
)

// BlockBuilderFactory represents a concrete BlockBuilderFactory implementation
type BlockBuilderFactory struct {
}

// CreateBlockBuilderFactory creates a new BlockBuilderFactory instance
func CreateBlockBuilderFactory() validated.BlockBuilderFactory {
	out := BlockBuilderFactory{}
	return &out
}

// Create creates a new BlockBuilder instance
func (fac *BlockBuilderFactory) Create() validated.BlockBuilder {
	out := createBlockBuilder()
	return out
}
