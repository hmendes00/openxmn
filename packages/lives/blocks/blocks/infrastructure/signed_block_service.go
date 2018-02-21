package infrastructure

import (
	"path/filepath"

	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	metadata "github.com/XMNBlockchain/core/packages/lives/metadata/domain"
	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
	stored_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/domain"
)

// SignedBlockService represents a concrete SignedBlockService implementation
type SignedBlockService struct {
	metaDataBuilderFactory  metadata.MetaDataBuilderFactory
	metaDataService         metadata.MetaDataService
	userSigService          users.SignatureService
	blkService              blocks.BlockService
	storedBlkBuilderFactory stored_blocks.SignedBlockBuilderFactory
}

// CreateSignedBlockService creates a new SignedBlockService instance
func CreateSignedBlockService(
	metaDataBuilderFactory metadata.MetaDataBuilderFactory,
	metaDataService metadata.MetaDataService,
	userSigService users.SignatureService,
	blkService blocks.BlockService,
	storedBlkBuilderFactory stored_blocks.SignedBlockBuilderFactory,
) blocks.SignedBlockService {
	out := SignedBlockService{
		metaDataBuilderFactory:  metaDataBuilderFactory,
		metaDataService:         metaDataService,
		userSigService:          userSigService,
		blkService:              blkService,
		storedBlkBuilderFactory: storedBlkBuilderFactory,
	}
	return &out
}

// Save saves a block instance
func (serv *SignedBlockService) Save(dirPath string, signedBlk blocks.SignedBlock) (stored_blocks.SignedBlock, error) {
	//build the metadata:
	id := signedBlk.GetID()
	ts := signedBlk.CreatedOn()
	met, metErr := serv.metaDataBuilderFactory.Create().Create().WithID(id).CreatedOn(ts).Now()
	if metErr != nil {
		return nil, metErr
	}

	//save the metadata:
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the signature:
	sig := signedBlk.GetSignature()
	storedSig, storedSigErr := serv.userSigService.Save(dirPath, sig)
	if storedSigErr != nil {
		return nil, storedSigErr
	}

	//save the block:
	blk := signedBlk.GetBlock()
	blkPath := filepath.Join(dirPath, "block")
	storedBlk, storedBlkErr := serv.blkService.Save(blkPath, blk)
	if storedBlkErr != nil {
		return nil, storedBlkErr
	}

	//build the stored block:
	storedSignedBlk, storedSignedBlkErr := serv.storedBlkBuilderFactory.Create().Create().WithBlock(storedBlk).WithMetaData(storedMet).WithSignature(storedSig).Now()
	if storedSignedBlkErr != nil {
		return nil, storedSignedBlkErr
	}

	return storedSignedBlk, nil
}
