package chained

import (
	validated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
	chained "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated/chained"
	concrete_validated "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/blocks/validated"
)

// Block represents a concrete block implementation
type Block struct {
	Met *MetaData                 `json:"metadata"`
	Blk *concrete_validated.Block `json:"block"`
}

func createBlock(met *MetaData, blk *concrete_validated.Block) chained.Block {
	out := Block{
		Met: met,
		Blk: blk,
	}

	return &out
}

// GetMetaData returns the MetaData
func (blk *Block) GetMetaData() chained.MetaData {
	return blk.Met
}

// GetBlock returns the Block
func (blk *Block) GetBlock() validated.Block {
	return blk.Blk
}
