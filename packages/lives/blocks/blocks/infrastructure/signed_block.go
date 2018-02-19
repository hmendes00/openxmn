package infrastructure

import (
	"time"

	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

// SignedBlock represents a concrete SignedBlock instance
type SignedBlock struct {
	ID   *uuid.UUID                `json:"id"`
	Blk  *Block                    `json:"block"`
	Sig  *concrete_users.Signature `json:"signature"`
	CrOn time.Time                 `json:"created_on"`
}

func createSignedBlock(id *uuid.UUID, blk *Block, sig *concrete_users.Signature, createdOn time.Time) blocks.SignedBlock {
	out := SignedBlock{
		ID:   id,
		Blk:  blk,
		Sig:  sig,
		CrOn: createdOn,
	}

	return &out
}

// GetID returns the ID
func (blk *SignedBlock) GetID() *uuid.UUID {
	return blk.ID
}

// GetBlock returns the Block
func (blk *SignedBlock) GetBlock() blocks.Block {
	return blk.Blk
}

// GetSignature returns the user Signature
func (blk *SignedBlock) GetSignature() users.Signature {
	return blk.Sig
}

// CreatedOn returns the creation time
func (blk *SignedBlock) CreatedOn() time.Time {
	return blk.CrOn
}
