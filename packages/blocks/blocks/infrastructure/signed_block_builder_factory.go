package infrastructure

import (
	blocks "github.com/XMNBlockchain/core/packages/blocks/blocks/domain"
)

// SignedBlockBuilderFactory represents a concrete SignedBlockBuilderFactory implementation
type SignedBlockBuilderFactory struct {
}

// CreateSignedBlockBuilderFactory creates a SignedBlockBuilderFactory instance
func CreateSignedBlockBuilderFactory() blocks.SignedBlockBuilderFactory {
	out := SignedBlockBuilderFactory{}
	return &out
}

// Create creates a new SignedBlockBuilder instance
func (fac *SignedBlockBuilderFactory) Create() blocks.SignedBlockBuilder {
	out := createSignedBlockBuilder()
	return out
}
