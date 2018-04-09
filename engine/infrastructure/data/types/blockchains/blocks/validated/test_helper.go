package validated

import (
	"strconv"
	"time"

	validated_blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
	concrete_stored_validated_blocks "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/blockchains/blocks/validated"
	concrete_blocks "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/blocks"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	uuid "github.com/satori/go.uuid"
)

// CreateBlockForTests creates a Block for tests
func CreateBlockForTests() *Block {
	//variables:
	id := uuid.NewV4()
	crOn := time.Now().UTC()
	signedBlk := concrete_blocks.CreateSignedBlockForTests()
	userSigs := concrete_users.CreateSignaturesForTests()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(crOn.UnixNano()))),
		signedBlk.GetMetaData().GetHashTree().GetHash().Get(),
		userSigs.GetMetaData().GetID().Bytes(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(crOn).Now()

	blk := createBlock(met.(*concrete_metadata.MetaData), signedBlk, userSigs)
	return blk.(*Block)
}

// CreateSignedBlockForTests creates a SignedBlock for tests
func CreateSignedBlockForTests() *SignedBlock {
	//variables:
	id := uuid.NewV4()
	crOn := time.Now().UTC()
	validatedBlk := CreateBlockForTests()
	sig := concrete_users.CreateSignatureForTests()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(crOn.UnixNano()))),
		validatedBlk.GetMetaData().GetHashTree().GetHash().Get(),
		sig.GetMetaData().GetID().Bytes(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(crOn).Now()

	blk := createSignedBlock(met.(*concrete_metadata.MetaData), validatedBlk, sig)
	return blk.(*SignedBlock)
}

// CreateBlockBuilderFactoryForTests creates a new BlockBuilderFactory for tests
func CreateBlockBuilderFactoryForTests() validated_blocks.BlockBuilderFactory {
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactoryForTests()
	out := CreateBlockBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	return out
}

// CreateBlockRepositoryForTests creates a new BlockRepository for tests
func CreateBlockRepositoryForTests() validated_blocks.BlockRepository {
	metaDataRepository := concrete_metadata.CreateRepositoryForTests()
	signedBlkRepository := concrete_blocks.CreateSignedBlockRepositoryForTests()
	userSigRepository := concrete_users.CreateSignaturesRepositoryForTests()
	validatedBlkBuilderFactory := CreateBlockBuilderFactoryForTests()
	out := CreateBlockRepository(metaDataRepository, signedBlkRepository, userSigRepository, validatedBlkBuilderFactory)
	return out
}

// CreateBlockServiceForTests creates a new BlockService for tests
func CreateBlockServiceForTests() validated_blocks.BlockService {
	metaDataService := concrete_metadata.CreateServiceForTests()
	signedBlkService := concrete_blocks.CreateSignedBlockServiceForTests()
	userSigService := concrete_users.CreateSignaturesServiceForTests()
	storedValidatedBlkBuilderFactory := concrete_stored_validated_blocks.CreateBlockBuilderFactoryForTests()
	out := CreateBlockService(metaDataService, signedBlkService, userSigService, storedValidatedBlkBuilderFactory)
	return out
}

// CreateSignedBlockBuilderFactoryForTests creates a new SignedBlockBuilderFactory for tests
func CreateSignedBlockBuilderFactoryForTests() validated_blocks.SignedBlockBuilderFactory {
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactoryForTests()
	out := CreateSignedBlockBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	return out
}

// CreateSignedBlockRepositoryForTests creates a new SignedBlockRepository for tests
func CreateSignedBlockRepositoryForTests() validated_blocks.SignedBlockRepository {
	metaDataRepository := concrete_metadata.CreateRepositoryForTests()
	validatedBlkRepository := CreateBlockRepositoryForTests()
	userSigRepository := concrete_users.CreateSignatureRepositoryForTests()
	signedBlkBuilderFactory := CreateSignedBlockBuilderFactoryForTests()
	out := CreateSignedBlockRepository(metaDataRepository, validatedBlkRepository, userSigRepository, signedBlkBuilderFactory)
	return out
}

// CreateSignedBlockServiceForTests creates a new SignedBlockService for tests
func CreateSignedBlockServiceForTests() validated_blocks.SignedBlockService {
	metaDataService := concrete_metadata.CreateServiceForTests()
	validatedBlkService := CreateBlockServiceForTests()
	userSigService := concrete_users.CreateSignatureServiceForTests()
	storedSignedValidatedBlkBuilderFactory := concrete_stored_validated_blocks.CreateSignedBlockBuilderFactoryForTests()
	out := CreateSignedBlockService(metaDataService, validatedBlkService, userSigService, storedSignedValidatedBlkBuilderFactory)
	return out
}
