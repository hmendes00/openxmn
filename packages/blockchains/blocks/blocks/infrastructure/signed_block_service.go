package infrastructure

import (
	"path/filepath"

	blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	stored_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/domain"
)

// SignedBlockService represents a concrete SignedBlockService implementation
type SignedBlockService struct {
	metaDataService         metadata.MetaDataService
	userSigService          users.SignatureService
	blkService              blocks.BlockService
	storedBlkBuilderFactory stored_blocks.SignedBlockBuilderFactory
}

// CreateSignedBlockService creates a new SignedBlockService instance
func CreateSignedBlockService(
	metaDataService metadata.MetaDataService,
	userSigService users.SignatureService,
	blkService blocks.BlockService,
	storedBlkBuilderFactory stored_blocks.SignedBlockBuilderFactory,
) blocks.SignedBlockService {
	out := SignedBlockService{
		metaDataService:         metaDataService,
		userSigService:          userSigService,
		blkService:              blkService,
		storedBlkBuilderFactory: storedBlkBuilderFactory,
	}
	return &out
}

// Save saves a block instance
func (serv *SignedBlockService) Save(dirPath string, signedBlk blocks.SignedBlock) (stored_blocks.SignedBlock, error) {
	//save the metadata:
	met := signedBlk.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the signature:
	sig := signedBlk.GetSignature()
	sigPath := filepath.Join(dirPath, "signature")
	storedSig, storedSigErr := serv.userSigService.Save(sigPath, sig)
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
