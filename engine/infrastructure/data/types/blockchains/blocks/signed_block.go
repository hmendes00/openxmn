package blocks

import (
	blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
)

// SignedBlock represents a concrete SignedBlock instance
type SignedBlock struct {
	Met *concrete_metadata.MetaData `json:"metadata"`
	Blk *Block                      `json:"block"`
	Sig *concrete_users.Signature   `json:"signature"`
}

func createSignedBlock(met *concrete_metadata.MetaData, blk *Block, sig *concrete_users.Signature) blocks.SignedBlock {
	out := SignedBlock{
		Met: met,
		Blk: blk,
		Sig: sig,
	}

	return &out
}

// GetMetaData returns the MetaData
func (blk *SignedBlock) GetMetaData() metadata.MetaData {
	return blk.Met
}

// GetBlock returns the Block
func (blk *SignedBlock) GetBlock() blocks.Block {
	return blk.Blk
}

// GetSignature returns the user Signature
func (blk *SignedBlock) GetSignature() users.Signature {
	return blk.Sig
}
