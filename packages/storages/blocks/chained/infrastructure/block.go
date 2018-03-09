package infrastructure

import (
	stored_chained_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/chained/domain"
	stored_validated_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/validated/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

type block struct {
	met stored_files.File
	blk stored_validated_blocks.Block
}

func createBlock(met stored_files.File, blk stored_validated_blocks.Block) stored_chained_blocks.Block {
	out := block{
		met: met,
		blk: blk,
	}

	return &out
}

// GetMetaData returns the MetaData file
func (blk *block) GetMetaData() stored_files.File {
	return blk.met
}

// GetBlock returns the stored validated block
func (blk *block) GetBlock() stored_validated_blocks.Block {
	return blk.blk
}
