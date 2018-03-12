package validated

import (
	stored_validated_block "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks/validated"
)

// CreateBlockBuilderFactoryForTests creates a BlockBuilderFactory for tests
func CreateBlockBuilderFactoryForTests() stored_validated_block.BlockBuilderFactory {
	out := CreateBlockBuilderFactory()
	return out
}
