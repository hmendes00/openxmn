package infrastructure

import (
	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	concrete_blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/infrastructure"
	validated "github.com/XMNBlockchain/core/packages/lives/blocks/validated/domain"
	hashtree "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	concrete_hashtree "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
)

// Block represents a concrete validated Block implementation
type Block struct {
	HT  *concrete_hashtree.HashTree  `json:"hashtree"`
	Blk *concrete_blocks.SignedBlock `json:"block"`
	LS  []*concrete_users.Signature  `json:"leader_signatures"`
}

func createBlock(ht *concrete_hashtree.HashTree, blk *concrete_blocks.SignedBlock, ls []*concrete_users.Signature) validated.Block {
	out := Block{
		HT:  ht,
		Blk: blk,
		LS:  ls,
	}

	return &out
}

// GetHashTree returns the HashTree
func (blk *Block) GetHashTree() hashtree.HashTree {
	return blk.HT
}

// GetBlock returns the Block
func (blk *Block) GetBlock() blocks.SignedBlock {
	return blk.Blk
}

// GetLeaderSignatures returns the leader Signatures
func (blk *Block) GetLeaderSignatures() []users.Signature {
	out := []users.Signature{}
	for _, oneLS := range blk.LS {
		out = append(out, oneLS)
	}

	return out
}
