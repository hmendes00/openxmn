package chained

import (
	"path/filepath"

	stored_validated_blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/blocks/validated"
	stored_chained_blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/blocks/validated/chained"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// BlockRepository represents a concrete stored chained BlockRepository implementation
type BlockRepository struct {
	fileRepository           stored_files.FileRepository
	validatedBlkRepository   stored_validated_blocks.BlockRepository
	chainedBlkBuilderFactory stored_chained_blocks.BlockBuilderFactory
}

// CreateBlockRepository creates a new BlockRepository instance
func CreateBlockRepository(fileRepository stored_files.FileRepository, validatedBlkRepository stored_validated_blocks.BlockRepository, chainedBlkBuilderFactory stored_chained_blocks.BlockBuilderFactory) stored_chained_blocks.BlockRepository {
	out := BlockRepository{
		fileRepository:           fileRepository,
		validatedBlkRepository:   validatedBlkRepository,
		chainedBlkBuilderFactory: chainedBlkBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a chained Block instance
func (rep *BlockRepository) Retrieve(dirPath string) (stored_chained_blocks.Block, error) {
	metPath := filepath.Join(dirPath, "metadata.json")
	met, metErr := rep.fileRepository.Retrieve(metPath)
	if metErr != nil {
		return nil, metErr
	}

	blkDirPath := filepath.Join(dirPath, "validated_block")
	blk, blkErr := rep.validatedBlkRepository.Retrieve(blkDirPath)
	if blkErr != nil {
		return nil, blkErr
	}

	out, outErr := rep.chainedBlkBuilderFactory.Create().Create().WithMetaData(met).WithBlock(blk).Now()
	if outErr != nil {
		return nil, outErr
	}

	return out, nil
}
