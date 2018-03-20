package validated

import (
	stored_validated_block "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/blocks/validated"
	concrete_stored_block "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/blockchains/blocks"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/blockchains/files"
	concrete_stored_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/blockchains/users"
)

// CreateBlockForTests creates a Block for tests
func CreateBlockForTests() *Block {
	met := concrete_stored_files.CreateFileForTests()
	blk := concrete_stored_block.CreateSignedBlockForTests()
	sigs := concrete_stored_users.CreateSignaturesForTests()
	out := createBlock(met, blk, sigs)
	return out.(*Block)
}

// CreateBlockBuilderFactoryForTests creates a BlockBuilderFactory for tests
func CreateBlockBuilderFactoryForTests() stored_validated_block.BlockBuilderFactory {
	out := CreateBlockBuilderFactory()
	return out
}

// CreateBlockRepositoryForTests creates a BlockRepository for tests
func CreateBlockRepositoryForTests() stored_validated_block.BlockRepository {
	fileRepository := concrete_stored_files.CreateFileRepositoryForTests()
	signedBlkRepository := concrete_stored_block.CreateSignedBlockRepositoryForTests()
	sigsRepository := concrete_stored_users.CreateSignaturesRepositoryForTests()
	validatedBlkBuilderFactory := CreateBlockBuilderFactoryForTests()
	out := CreateBlockRepository(fileRepository, signedBlkRepository, sigsRepository, validatedBlkBuilderFactory)
	return out
}
