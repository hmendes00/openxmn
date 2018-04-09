package validated

import (
	stored_validated_block "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/blocks/validated"
	concrete_stored_block "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/blockchains/blocks"
	concrete_stored_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/blockchains/users"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
)

// CreateBlockForTests creates a Block for tests
func CreateBlockForTests() *Block {
	met := concrete_stored_files.CreateFileForTests()
	blk := concrete_stored_block.CreateSignedBlockForTests()
	sigs := concrete_stored_users.CreateSignaturesForTests()
	out := createBlock(met, blk, sigs)
	return out.(*Block)
}

// CreateSignedBlockForTests creates a SignedBlock for tests
func CreateSignedBlockForTests() *SignedBlock {
	met := concrete_stored_files.CreateFileForTests()
	blk := CreateBlockForTests()
	sig := concrete_stored_users.CreateSignatureForTests()
	out := createSignedBlock(met, blk, sig)
	return out.(*SignedBlock)
}

// CreateBlockBuilderFactoryForTests creates a BlockBuilderFactory for tests
func CreateBlockBuilderFactoryForTests() stored_validated_block.BlockBuilderFactory {
	out := CreateBlockBuilderFactory()
	return out
}

// CreateSignedBlockBuilderFactoryForTests creates a SignedBlockBuilderFactory for tests
func CreateSignedBlockBuilderFactoryForTests() stored_validated_block.SignedBlockBuilderFactory {
	out := CreateSignedBlockBuilderFactory()
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
