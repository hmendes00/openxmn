package blocks

import (
	stored_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks"
	conrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/files"
	concrete_stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/transactions/signed/aggregated"
	concrete_stored_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/users"
)

// CreateBlockForTests creates a Block for tests
func CreateBlockForTests() *Block {
	met := conrete_stored_files.CreateFileForTests()
	trs := []*concrete_stored_aggregated_transactions.SignedTransactions{
		concrete_stored_aggregated_transactions.CreateSignedTransactionsForTests(),
		concrete_stored_aggregated_transactions.CreateSignedTransactionsForTests(),
		concrete_stored_aggregated_transactions.CreateSignedTransactionsForTests(),
	}

	out := createBlock(met, trs)
	return out.(*Block)
}

// CreateSignedBlockForTests creates a SignedBlock for tests
func CreateSignedBlockForTests() *SignedBlock {
	met := conrete_stored_files.CreateFileForTests()
	sig := concrete_stored_users.CreateSignatureForTests()
	blk := CreateBlockForTests()

	out := createSignedBlock(met, sig, blk)
	return out.(*SignedBlock)
}

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

// CreateBlockRepositoryForTests creates a BlockRepository for tests
func CreateBlockRepositoryForTests() stored_blocks.BlockRepository {
	fileRepository := conrete_stored_files.CreateFileRepositoryForTests()
	aggrTransRepository := concrete_stored_aggregated_transactions.CreateSignedTransactionsRepositoryForTests()
	blkBuilderFactory := CreateBlockBuilderFactoryForTests()
	out := CreateBlockRepository(fileRepository, aggrTransRepository, blkBuilderFactory)
	return out
}

// CreateSignedBlockRepositoryForTests creates a SignedBlockRepository for tests
func CreateSignedBlockRepositoryForTests() stored_blocks.SignedBlockRepository {
	fileRepository := conrete_stored_files.CreateFileRepositoryForTests()
	sigRepository := concrete_stored_users.CreateSignatureRepositoryForTests()
	blkRepository := CreateBlockRepositoryForTests()
	signedBlkBuilderFactory := CreateSignedBlockBuilderFactoryForTests()
	out := CreateSignedBlockRepository(fileRepository, sigRepository, blkRepository, signedBlkBuilderFactory)
	return out
}
