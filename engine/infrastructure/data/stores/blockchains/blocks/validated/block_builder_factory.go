package validated

import (
	stored_validated_block "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/blocks/validated"
)

// BlockBuilderFactory represents a concrete BlockBuilderFactory implementation
type BlockBuilderFactory struct {
}

// CreateBlockBuilderFactory creates a new BlockBuilderFactory instance
func CreateBlockBuilderFactory() stored_validated_block.BlockBuilderFactory {
	out := BlockBuilderFactory{}
	return &out
}

// Create creates a new BlockBuilder instance
func (fac *BlockBuilderFactory) Create() stored_validated_block.BlockBuilder {
	out := createBlockBuilder()
	return out
}
