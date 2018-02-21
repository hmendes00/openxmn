package infrastructure

import (
	"time"

	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	concrete_blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/infrastructure"
	validated "github.com/XMNBlockchain/core/packages/lives/blocks/validated/domain"
	hashtree "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
	concrete_hashtree "github.com/XMNBlockchain/core/packages/lives/hashtrees/infrastructure"
	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/lives/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// Block represents a concrete validated Block implementation
type Block struct {
	ID   *uuid.UUID                   `json:"id"`
	HT   *concrete_hashtree.HashTree  `json:"hashtree"`
	Blk  *concrete_blocks.SignedBlock `json:"block"`
	LS   []*concrete_users.Signature  `json:"leader_signatures"`
	CrOn time.Time                    `json:"created_on"`
}

func createBlock(id *uuid.UUID, ht *concrete_hashtree.HashTree, blk *concrete_blocks.SignedBlock, ls []*concrete_users.Signature, createdOn time.Time) validated.Block {
	out := Block{
		ID:   id,
		HT:   ht,
		Blk:  blk,
		LS:   ls,
		CrOn: createdOn,
	}

	return &out
}

// GetID returns the ID
func (blk *Block) GetID() *uuid.UUID {
	return blk.ID
}

// GetHashTree returns the HashTree
func (blk *Block) GetHashTree() hashtree.HashTree {
	return blk.HT
}

// GetBlock returns the Block
func (blk *Block) GetBlock() blocks.SignedBlock {
	return blk.Blk
}

// GetSignatures returns the leader Signatures
func (blk *Block) GetSignatures() []users.Signature {
	out := []users.Signature{}
	for _, oneLS := range blk.LS {
		out = append(out, oneLS)
	}

	return out
}

// CreatedOn returns the creation time
func (blk *Block) CreatedOn() time.Time {
	return blk.CrOn
}
