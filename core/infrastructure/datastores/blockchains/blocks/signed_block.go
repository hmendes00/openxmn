package blocks

import (
	blocks "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/blocks"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/metadata"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/users"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/metadata"
	concrete_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/users"
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
