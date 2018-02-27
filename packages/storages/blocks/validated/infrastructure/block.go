package infrastructure

import (
	stored_block "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/domain"
	stored_validated_block "github.com/XMNBlockchain/core/packages/storages/blocks/validated/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

type block struct {
	metaData stored_files.File
	blk      stored_block.SignedBlock
	sigs     []stored_files.File
}

func createBlock(metaData stored_files.File, blk stored_block.SignedBlock, sigs []stored_files.File) stored_validated_block.Block {
	out := block{
		metaData: metaData,
		blk:      blk,
		sigs:     sigs,
	}

	return &out
}

// GetMetaData returns the metadata file
func (blk *block) GetMetaData() stored_files.File {
	return blk.metaData
}

// GetBlock returns the stored block
func (blk *block) GetBlock() stored_block.SignedBlock {
	return blk.blk
}

// GetSignatures returns the stored signatures
func (blk *block) GetSignatures() []stored_files.File {
	return blk.sigs
}
