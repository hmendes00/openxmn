package chained

import (
	"path/filepath"

	validated "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/blocks/validated"
	chained "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/blocks/validated/chained"
)

// BlockRepository represents a concrete BlockRepository implementation
type BlockRepository struct {
	metaDataRepository chained.MetaDataRepository
	valBlkRepository   validated.BlockRepository
	blkBuilderFactory  chained.BlockBuilderFactory
}

// CreateBlockRepository creates a new BlockRepository instance
func CreateBlockRepository(metaDataRepository chained.MetaDataRepository, valBlkRepository validated.BlockRepository, blkBuilderFactory chained.BlockBuilderFactory) chained.BlockRepository {
	out := BlockRepository{
		metaDataRepository: metaDataRepository,
		valBlkRepository:   valBlkRepository,
		blkBuilderFactory:  blkBuilderFactory,
	}

	return &out
}

// Retrieve retrieve a chained block
func (rep *BlockRepository) Retrieve(dirPath string) (chained.Block, error) {
	//retrieve the metadata:
	met, metErr := rep.metaDataRepository.Retrieve(dirPath)
	if metErr != nil {
		return nil, metErr
	}

	//retrieve the validated block:
	blkDirPath := filepath.Join(dirPath, "validated_block")
	valBlk, valBlkErr := rep.valBlkRepository.Retrieve(blkDirPath)
	if valBlkErr != nil {
		return nil, valBlkErr
	}

	//build the chained block:
	out, outErr := rep.blkBuilderFactory.Create().Create().WithMetaData(met).WithBlock(valBlk).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
