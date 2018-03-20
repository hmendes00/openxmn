package chained

import (
	validated "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/blocks/validated"
	chained "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/blocks/validated/chained"
	concrete_validated "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/blocks/validated"
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
