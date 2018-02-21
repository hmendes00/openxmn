package infrastructure

import (
	stored_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/domain"
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
