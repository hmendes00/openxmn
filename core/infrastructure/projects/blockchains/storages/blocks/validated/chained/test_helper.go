package chained

import (
	stored_chained_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks/validated/chained"
	concrete_stored_validated_blocks "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/blocks/validated"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/files"
)

// CreateBlockForTests creates a Block for tests
func CreateBlockForTests() *Block {
	met := concrete_stored_files.CreateFileForTests()
	blk := concrete_stored_validated_blocks.CreateBlockForTests()
	out := createBlock(met, blk)
	return out.(*Block)
}

// CreateBlockBuilderFactoryForTests creates a BlockBuilderFactory for tests
func CreateBlockBuilderFactoryForTests() stored_chained_blocks.BlockBuilderFactory {
	out := CreateBlockBuilderFactory()
	return out
}
