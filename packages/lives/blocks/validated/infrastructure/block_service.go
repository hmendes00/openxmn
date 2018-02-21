package infrastructure

import (
	"path/filepath"

	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	validated "github.com/XMNBlockchain/core/packages/lives/blocks/validated/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/lives/metadata/domain"
	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
	stored_validated_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/validated/domain"
)

// BlockService represents a concrete BlockService implementation
type BlockService struct {
	metaDataBuilderFactory           metadata.MetaDataBuilderFactory
	metaDataService                  metadata.MetaDataService
	htService                        hashtrees.HashTreeService
	signedBlkService                 blocks.SignedBlockService
	userSigService                   users.SignatureService
	storedValidatedBlkBuilderFactory stored_validated_blocks.BlockBuilderFactory
}

// CreateBlockService creates a new BlockService instance
func CreateBlockService(
	metaDataBuilderFactory metadata.MetaDataBuilderFactory,
	metaDataService metadata.MetaDataService,
	htService hashtrees.HashTreeService,
	signedBlkService blocks.SignedBlockService,
	userSigService users.SignatureService,
	storedValidatedBlkBuilderFactory stored_validated_blocks.BlockBuilderFactory,
) validated.BlockService {
	out := BlockService{
		metaDataBuilderFactory:           metaDataBuilderFactory,
		metaDataService:                  metaDataService,
		htService:                        htService,
		signedBlkService:                 signedBlkService,
		userSigService:                   userSigService,
		storedValidatedBlkBuilderFactory: storedValidatedBlkBuilderFactory,
	}
	return &out
}

// Save saves a Block
func (serv *BlockService) Save(dirPath string, validatedBlk validated.Block) (stored_validated_blocks.Block, error) {
	//build the metadata:
	id := validatedBlk.GetID()
	ts := validatedBlk.CreatedOn()
	met, metErr := serv.metaDataBuilderFactory.Create().Create().WithID(id).CreatedOn(ts).Now()
	if metErr != nil {
		return nil, metErr
	}

	//save the metadata:
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the hashtree:
	ht := validatedBlk.GetHashTree()
	storedHt, storedHtErr := serv.htService.Save(dirPath, ht)
	if storedHtErr != nil {
		return nil, storedHtErr
	}

	//save the block:
	blk := validatedBlk.GetBlock()
	blkPath := filepath.Join(dirPath, "signed_block")
	storedBlk, storedBlkErr := serv.signedBlkService.Save(blkPath, blk)
	if storedBlkErr != nil {
		return nil, storedBlkErr
	}

	//save the user signatures:
	sigs := validatedBlk.GetSignatures()
	sigsPath := filepath.Join(dirPath, "user_signatures")
	storedSigs, storedSigsErr := serv.userSigService.SaveAll(sigsPath, sigs)
	if storedSigsErr != nil {
		return nil, storedSigsErr
	}

	//build the stored validated block:
	storedValidatedBlk, storedValidatedBlkErr := serv.storedValidatedBlkBuilderFactory.Create().Create().WithMetaData(storedMet).WithHashTree(storedHt).WithBlock(storedBlk).WithSignatures(storedSigs).Now()
	if storedValidatedBlkErr != nil {
		return nil, storedValidatedBlkErr
	}

	return storedValidatedBlk, nil
}
