package validated

import (
	validated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
)

// SignedBlock represents a concrete SignedBlock implementation
type SignedBlock struct {
	Met *concrete_metadata.MetaData `json:"metadata"`
	Blk *Block                      `json:"block"`
	Sig *concrete_users.Signature   `json:"signature"`
}

func createSignedBlock(met *concrete_metadata.MetaData, blk *Block, sig *concrete_users.Signature) validated.SignedBlock {
	out := SignedBlock{
		Met: met,
		Blk: blk,
		Sig: sig,
	}

	return &out
}

// GetMetaData returns the metadata
func (blk *SignedBlock) GetMetaData() metadata.MetaData {
	return blk.Met
}

// GetBlock returns the block
func (blk *SignedBlock) GetBlock() validated.Block {
	return blk.Blk
}

// GetSignature returns the signature
func (blk *SignedBlock) GetSignature() users.Signature {
	return blk.Sig
}
