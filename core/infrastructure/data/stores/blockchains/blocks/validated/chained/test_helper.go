package chained

import (
	stored_chained_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/blocks/validated/chained"
	concrete_stored_validated_blocks "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/blockchains/blocks/validated"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/blockchains/files"
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

// CreateBlockRepositoryForTests creates a BlockRepository for tests
func CreateBlockRepositoryForTests() stored_chained_blocks.BlockRepository {
	fileRepository := concrete_stored_files.CreateFileRepositoryForTests()
	validatedBlkRepository := concrete_stored_validated_blocks.CreateBlockRepositoryForTests()
	chainedBlkBuilderFactory := CreateBlockBuilderFactoryForTests()
	out := CreateBlockRepository(fileRepository, validatedBlkRepository, chainedBlkBuilderFactory)
	return out
}
