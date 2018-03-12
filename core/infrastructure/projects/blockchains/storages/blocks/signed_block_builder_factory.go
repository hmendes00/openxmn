package blocks

import (
	stored_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks"
)

// SignedBlockBuilderFactory represents a concrete SignedBlockBuilderFactory implementation
type SignedBlockBuilderFactory struct {
}

// CreateSignedBlockBuilderFactory creates a new SignedBlockBuilderFactory instance
func CreateSignedBlockBuilderFactory() stored_blocks.SignedBlockBuilderFactory {
	out := SignedBlockBuilderFactory{}
	return &out
}

// Create creates a new SignedBlockBuilder instance
func (fac *SignedBlockBuilderFactory) Create() stored_blocks.SignedBlockBuilder {
	out := createSignedBlockBuilder()
	return out
}
