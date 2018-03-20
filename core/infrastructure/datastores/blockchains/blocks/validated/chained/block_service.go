package chained

import (
	"path/filepath"

	validated "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/blocks/validated"
	chained "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/blocks/validated/chained"
	stored_chained_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/blocks/validated/chained"
)

// BlockService represents a concrete BlockService implementation
type BlockService struct {
	metaDataService                chained.MetaDataService
	blkService                     validated.BlockService
	storedChainedBlkBuilderFactory stored_chained_blocks.BlockBuilderFactory
}

// CreateBlockService creates a new BlockService instance
func CreateBlockService(metaDataService chained.MetaDataService, blkService validated.BlockService, storedChainedBlkBuilderFactory stored_chained_blocks.BlockBuilderFactory) chained.BlockService {
	out := BlockService{
		metaDataService:                metaDataService,
		blkService:                     blkService,
		storedChainedBlkBuilderFactory: storedChainedBlkBuilderFactory,
	}

	return &out
}

// Save saves a chained Block
func (serv *BlockService) Save(dirPath string, blk chained.Block) (stored_chained_blocks.Block, error) {
	//save the metadata:
	met := blk.GetMetaData()
	storedMet, storedMetErr := serv.metaDataService.Save(dirPath, met)
	if storedMetErr != nil {
		return nil, storedMetErr
	}

	//save the validated block:
	valBlk := blk.GetBlock()
	blkDirPath := filepath.Join(dirPath, "validated_block")
	storedValidatedBlk, storedValidatedBlkErr := serv.blkService.Save(blkDirPath, valBlk)
	if storedValidatedBlkErr != nil {
		return nil, storedValidatedBlkErr
	}

	//build the stored chained blk:
	out, outErr := serv.storedChainedBlkBuilderFactory.Create().Create().WithMetaData(storedMet).WithBlock(storedValidatedBlk).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
