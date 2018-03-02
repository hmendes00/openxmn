package infrastructure

import (
	chained "github.com/XMNBlockchain/core/packages/blockchains/blocks/chained/domain"
	validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/domain"
	concrete_validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/infrastructure"
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
)

// Block represents a concrete block implementation
type Block struct {
	HT  *concrete_hashtrees.HashTree `json:"hashtree"`
	Met *MetaData                    `json:"metadata"`
	Blk *concrete_validated.Block    `json:"block"`
}

func createBlock(ht *concrete_hashtrees.HashTree, met *MetaData, blk *concrete_validated.Block) chained.Block {
	out := Block{
		HT:  ht,
		Met: met,
		Blk: blk,
	}

	return &out
}

// GetHashTree returns the HashTree
func (blk *Block) GetHashTree() hashtrees.HashTree {
	return blk.HT
}

// GetMetaData returns the MetaData
func (blk *Block) GetMetaData() chained.MetaData {
	return blk.Met
}

// GetBlock returns the Block
func (blk *Block) GetBlock() validated.Block {
	return blk.Blk
}
