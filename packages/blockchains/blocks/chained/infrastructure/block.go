package infrastructure

import (
	chained "github.com/XMNBlockchain/core/packages/blockchains/blocks/chained/domain"
	validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/domain"
	concrete_validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/infrastructure"
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
