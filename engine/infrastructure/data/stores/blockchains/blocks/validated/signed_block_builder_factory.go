package validated

import (
	stored_validated_block "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/blocks/validated"
)

// SignedBlockBuilderFactory represents a concrete SignedBlockBuilderFactory implementation
type SignedBlockBuilderFactory struct {
}

// CreateSignedBlockBuilderFactory creates a new SignedBlockBuilderFactory instance
func CreateSignedBlockBuilderFactory() stored_validated_block.SignedBlockBuilderFactory {
	out := SignedBlockBuilderFactory{}
	return &out
}

// Create creates a new SignedBlockBuilder instance
func (fac *SignedBlockBuilderFactory) Create() stored_validated_block.SignedBlockBuilder {
	out := createSignedBlockBuilder()
	return out
}
