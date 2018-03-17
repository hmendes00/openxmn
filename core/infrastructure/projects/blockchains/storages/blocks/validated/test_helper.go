package validated

import (
	stored_validated_block "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks/validated"
	concrete_stored_block "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/blocks"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/files"
	concrete_stored_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/users"
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
