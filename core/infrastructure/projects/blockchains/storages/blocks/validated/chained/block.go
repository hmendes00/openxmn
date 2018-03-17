package chained

import (
	stored_validated_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks/validated"
	stored_chained_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks/validated/chained"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	concrete_stored_validated_blocks "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/blocks/validated"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/files"
)

// Block represents a concrete stored chained block implementation
type Block struct {
	Met *concrete_stored_files.File             `json:"metadata"`
	Blk *concrete_stored_validated_blocks.Block `json:"validated_block"`
}

func createBlock(met *concrete_stored_files.File, blk *concrete_stored_validated_blocks.Block) stored_chained_blocks.Block {
	out := Block{
		Met: met,
		Blk: blk,
	}

	return &out
}

// GetMetaData returns the MetaData file
func (blk *Block) GetMetaData() stored_files.File {
	return blk.Met
}

// GetBlock returns the stored validated block
func (blk *Block) GetBlock() stored_validated_blocks.Block {
	return blk.Blk
}
