package validated

import (
	"strconv"
	"time"

	validated_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/blocks/validated"
	concrete_blocks "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/blocks"
	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/metadata"
	concrete_stored_validated_blocks "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/blockchains/blocks/validated"
	concrete_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/users"
	uuid "github.com/satori/go.uuid"
)

// CreateBlockForTests creates a Block for tests
func CreateBlockForTests() *Block {

	concrete_users.CreateSignatureForTests()
	//variables:
	id := uuid.NewV4()
	crOn := time.Now().UTC()
	signedBlk := concrete_blocks.CreateSignedBlockForTests()
	userSigs := concrete_users.CreateSignaturesForTests()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(crOn.UnixNano()))),
		signedBlk.GetMetaData().GetHashTree().GetHash().Get(),
		userSigs.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(crOn).Now()

	blk := createBlock(met.(*concrete_metadata.MetaData), signedBlk, userSigs)
	return blk.(*Block)
}

// CreateBlockBuilderFactoryForTests creates a new BlockBuilderFactory for tests
func CreateBlockBuilderFactoryForTests() validated_blocks.BlockBuilderFactory {
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactoryForTests()
	out := CreateBlockBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	return out
}

// CreateBlockRepositoryForTests creates a new BlockRepository for tests
func CreateBlockRepositoryForTests() validated_blocks.BlockRepository {
	metaDataRepository := concrete_metadata.CreateMetaDataRepositoryForTests()
	signedBlkRepository := concrete_blocks.CreateSignedBlockRepositoryForTests()
	userSigRepository := concrete_users.CreateSignaturesRepositoryForTests()
	validatedBlkBuilderFactory := CreateBlockBuilderFactoryForTests()
	out := CreateBlockRepository(metaDataRepository, signedBlkRepository, userSigRepository, validatedBlkBuilderFactory)
	return out
}

// CreateBlockServiceForTests creates a new BlockService for tests
func CreateBlockServiceForTests() validated_blocks.BlockService {
	metaDataService := concrete_metadata.CreateMetaDataServiceForTests()
	signedBlkService := concrete_blocks.CreateSignedBlockServiceForTests()
	userSigService := concrete_users.CreateSignaturesServiceForTests()
	storedValidatedBlkBuilderFactory := concrete_stored_validated_blocks.CreateBlockBuilderFactoryForTests()
	out := CreateBlockService(metaDataService, signedBlkService, userSigService, storedValidatedBlkBuilderFactory)
	return out
}
