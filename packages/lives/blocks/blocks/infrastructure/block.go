package infrastructure

import (
	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	hashtree "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	concrete_hashtree "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
	concrete_aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/infrastructure"
)

// Block represents a concrete Block implementation
type Block struct {
	HT          *concrete_hashtree.HashTree               `json:"hashtree"`
	Trs         []*concrete_aggregated.SignedTransactions `json:"transactions"`
	NeededKarma int                                       `json:"needed_karma"`
}

func createBlock(ht *concrete_hashtree.HashTree, trs []*concrete_aggregated.SignedTransactions, neededKarma int) blocks.Block {
	out := Block{
		HT:          ht,
		Trs:         trs,
		NeededKarma: neededKarma,
	}

	return &out
}

// GetHashTree returns the HashTree
func (blk *Block) GetHashTree() hashtree.HashTree {
	return blk.HT
}

// GetNeededKarma returns the needed karma to validate this block
func (blk *Block) GetNeededKarma() int {
	return blk.NeededKarma
}

// GetTransactions returns the Signed Transactions
func (blk *Block) GetTransactions() []aggregated.SignedTransactions {
	out := []aggregated.SignedTransactions{}
	for _, oneSignedTrs := range blk.Trs {
		out = append(out, oneSignedTrs)
	}

	return out
}
