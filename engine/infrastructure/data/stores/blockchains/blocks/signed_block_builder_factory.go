package blocks

import (
	stored_blocks "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/blockchains/blocks"
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
