package chained

import (
	stored_validated_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks/validated"
	stored_chained_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks/validated/chained"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
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
