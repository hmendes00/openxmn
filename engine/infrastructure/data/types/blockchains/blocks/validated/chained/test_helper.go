package chained

import (
	"strconv"
	"time"

	chained "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/blocks/validated/chained"
	concrete_validated "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/types/blockchains/blocks/validated"
	concrete_files "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/types/files"
	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/types/hashtrees"
	concrete_stored_chained "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/stores/blockchains/blocks/validated/chained"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/stores/files"
	uuid "github.com/satori/go.uuid"
)

// CreateMetaDataForTests creates a Metadata for tests
func CreateMetaDataForTests() *MetaData {
	//variables:
	id := uuid.NewV4()
	prevID := uuid.NewV4()
	cr := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(cr.UnixNano()))),
		prevID.Bytes(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()

	met := createMetaData(&id, ht.(*concrete_hashtrees.HashTree), &prevID, cr)
	return met.(*MetaData)
}

// CreateBlockForTests creates a Block for tests
func CreateBlockForTests() *Block {
	//variables:
	id := uuid.NewV4()
	prevID := uuid.NewV4()
	cr := time.Now().UTC()
	valBlk := concrete_validated.CreateBlockForTests()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(cr.UnixNano()))),
		prevID.Bytes(),
		valBlk.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()

	met := createMetaData(&id, ht.(*concrete_hashtrees.HashTree), &prevID, cr)
	chainedBlk := createBlock(met.(*MetaData), valBlk)
	return chainedBlk.(*Block)
}

// CreateMetaDataBuilderFactoryForTests creates a new MetaDataBuilderFactory for tests
func CreateMetaDataBuilderFactoryForTests() chained.MetaDataBuilderFactory {
	out := CreateMetaDataBuilderFactory()
	return out
}

// CreateMetaDataRepositoryForTests creates a new MetaDataRepository for tests
func CreateMetaDataRepositoryForTests() chained.MetaDataRepository {
	fileRepository := concrete_files.CreateFileRepositoryForTests()
	out := CreateMetaDataRepository(fileRepository)
	return out
}

// CreateMetaDataServiceForTests creates a new MetaDataService for tests
func CreateMetaDataServiceForTests() chained.MetaDataService {
	fileBuilderFactory := concrete_files.CreateFileBuilderFactoryForTests()
	fileService := concrete_files.CreateFileServiceForTests()
	storedFileBuilderFactory := concrete_stored_files.CreateFileBuilderFactoryForTests()
	out := CreateMetaDataService(fileBuilderFactory, fileService, storedFileBuilderFactory)
	return out
}

// CreateBlockBuilderFactoryForTests creates a new BlockBuilderFactory for tests
func CreateBlockBuilderFactoryForTests() chained.BlockBuilderFactory {
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	metaDataBuilderFactory := CreateMetaDataBuilderFactoryForTests()
	out := CreateBlockBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	return out
}

// CreateBlockRepositoryForTests creates a new BlockRepository for tests
func CreateBlockRepositoryForTests() chained.BlockRepository {
	metaDataRepository := CreateMetaDataRepositoryForTests()
	valBlkRepository := concrete_validated.CreateBlockRepositoryForTests()
	blkBuilderFactory := CreateBlockBuilderFactoryForTests()
	out := CreateBlockRepository(metaDataRepository, valBlkRepository, blkBuilderFactory)
	return out
}

// CreateBlockServiceForTests creates a new BlockService for tests
func CreateBlockServiceForTests() chained.BlockService {
	metaDataService := CreateMetaDataServiceForTests()
	blkService := concrete_validated.CreateBlockServiceForTests()
	storedChainedBlkBuilderFactory := concrete_stored_chained.CreateBlockBuilderFactoryForTests()
	out := CreateBlockService(metaDataService, blkService, storedChainedBlkBuilderFactory)
	return out
}
