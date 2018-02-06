package infrastructure

import (
	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
)

// SignedBlock represents a concrete SignedBlock instance
type SignedBlock struct {
	Blk *Block                    `json:"block"`
	Sig *concrete_users.Signature `json:"signature"`
}

func createSignedBlock(blk *Block, sig *concrete_users.Signature) blocks.SignedBlock {
	out := SignedBlock{
		Blk: blk,
		Sig: sig,
	}

	return &out
}

// GetBlock returns the Block
func (blk *SignedBlock) GetBlock() blocks.Block {
	return blk.Blk
}

// GetSignature returns the user Signature
func (blk *SignedBlock) GetSignature() users.Signature {
	return blk.Sig
}
