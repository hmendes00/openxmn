package blocks

import (
	"strconv"
	"time"

	blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/blocks"
	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/metadata"
	concrete_stored_blocks "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/blocks"
	concrete_aggregated "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/transactions/signed/aggregated"
	concrete_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/users"
	uuid "github.com/satori/go.uuid"
)

// CreateBlockForTests creates a Block for tests
func CreateBlockForTests() *Block {
	//variables:
	id := uuid.NewV4()
	cr := time.Now().UTC()
	trs := []*concrete_aggregated.SignedTransactions{
		concrete_aggregated.CreateSignedTransactionsForTests(),
		concrete_aggregated.CreateSignedTransactionsForTests(),
		concrete_aggregated.CreateSignedTransactionsForTests(),
	}

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(cr.UnixNano()))),
	}

	for _, oneTrs := range trs {
		blocks = append(blocks, oneTrs.GetMetaData().GetHashTree().GetHash().Get())
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(cr).Now()

	blk := createBlock(met.(*concrete_metadata.MetaData), trs)
	return blk.(*Block)
}

// CreateSignedBlockForTests creates a SignedBlock for tests
func CreateSignedBlockForTests() *SignedBlock {
	//variables:
	id := uuid.NewV4()
	blk := CreateBlockForTests()
	sig := concrete_users.CreateSignatureForTests()
	crOn := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(crOn.UnixNano()))),
		blk.GetMetaData().GetHashTree().GetHash().Get(),
		sig.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(crOn).Now()

	signedBlk := createSignedBlock(met.(*concrete_metadata.MetaData), blk, sig)
	return signedBlk.(*SignedBlock)
}

// CreateBlockBuilderFactoryForTests creates a new BlockBuilderFactory for tests
func CreateBlockBuilderFactoryForTests() blocks.BlockBuilderFactory {
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactoryForTests()
	out := CreateBlockBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	return out
}

// CreateBlockRepositoryForTests creates a new BlockRepository for tests
func CreateBlockRepositoryForTests() blocks.BlockRepository {
	metaDataRepository := concrete_metadata.CreateMetaDataRepositoryForTests()
	signedTrsRepository := concrete_aggregated.CreateSignedTransactionsRepositoryForTests()
	blkBuilderFactory := CreateBlockBuilderFactoryForTests()
	out := CreateBlockRepository(metaDataRepository, signedTrsRepository, blkBuilderFactory)
	return out
}

// CreateBlockServiceForTests creates a new BlockService for tests
func CreateBlockServiceForTests() blocks.BlockService {
	metaDataService := concrete_metadata.CreateMetaDataServiceForTests()
	signedTrsService := concrete_aggregated.CreateSignedTransactionsServiceForTests()
	storedBlkBuilderFactory := concrete_stored_blocks.CreateBlockBuilderFactoryForTests()
	out := CreateBlockService(metaDataService, signedTrsService, storedBlkBuilderFactory)
	return out
}

// CreateSignedBlockBuilderFactoryForTests creates a new SignedBlockBuilderFactory for tests
func CreateSignedBlockBuilderFactoryForTests() blocks.SignedBlockBuilderFactory {
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactoryForTests()
	out := CreateSignedBlockBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	return out
}

// CreateSignedBlockRepositoryForTests creates a new SignedBlockRepository for tests
func CreateSignedBlockRepositoryForTests() blocks.SignedBlockRepository {
	metaDataRepository := concrete_metadata.CreateMetaDataRepositoryForTests()
	userSigRepository := concrete_users.CreateSignatureRepositoryForTests()
	blkRepository := CreateBlockRepositoryForTests()
	signedBlkBuilderFactory := CreateSignedBlockBuilderFactoryForTests()
	out := CreateSignedBlockRepository(metaDataRepository, userSigRepository, blkRepository, signedBlkBuilderFactory)
	return out
}

// CreateSignedBlockServiceForTests creates a new SignedBlockService for tests
func CreateSignedBlockServiceForTests() blocks.SignedBlockService {
	metaDataService := concrete_metadata.CreateMetaDataServiceForTests()
	userSigService := concrete_users.CreateSignatureServiceForTests()
	blkService := CreateBlockServiceForTests()
	storedBlkBuilderFactory := concrete_stored_blocks.CreateSignedBlockBuilderFactoryForTests()
	out := CreateSignedBlockService(metaDataService, userSigService, blkService, storedBlkBuilderFactory)
	return out
}
