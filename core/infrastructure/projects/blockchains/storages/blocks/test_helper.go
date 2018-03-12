package blocks

import (
	stored_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks"
)

// CreateBlockBuilderFactoryForTests creates a BlockBuilderFactory for tests
func CreateBlockBuilderFactoryForTests() stored_blocks.BlockBuilderFactory {
	out := CreateBlockBuilderFactory()
	return out
}

// CreateSignedBlockBuilderFactoryForTests creates a SignedBlockBuilderFactory for tests
func CreateSignedBlockBuilderFactoryForTests() stored_blocks.SignedBlockBuilderFactory {
	out := CreateSignedBlockBuilderFactory()
	return out
}
