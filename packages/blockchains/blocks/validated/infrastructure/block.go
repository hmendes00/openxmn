package infrastructure

import (
	"time"

	blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/domain"
	concrete_blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/infrastructure"
	validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// Block represents a concrete validated Block implementation
type Block struct {
	ID   *uuid.UUID                   `json:"id"`
	Blk  *concrete_blocks.SignedBlock `json:"block"`
	LS   []*concrete_users.Signature  `json:"leader_signatures"`
	CrOn time.Time                    `json:"created_on"`
}

func createBlock(id *uuid.UUID, blk *concrete_blocks.SignedBlock, ls []*concrete_users.Signature, createdOn time.Time) validated.Block {
	out := Block{
		ID:   id,
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
