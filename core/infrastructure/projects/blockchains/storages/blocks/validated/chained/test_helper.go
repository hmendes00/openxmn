package chained

import (
	tored_chained_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks/validated/chained"
)

// CreateBlockBuilderFactoryForTests creates a BlockBuilderFactory for tests
func CreateBlockBuilderFactoryForTests() tored_chained_blocks.BlockBuilderFactory {
	out := CreateBlockBuilderFactory()
	return out
}
