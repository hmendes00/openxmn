package validated

import (
	stored_validated_block "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/blocks/validated"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/users"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	concrete_stored_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/users"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
)

// SignedBlock represents a concrete SignedBlock implementation
type SignedBlock struct {
	Met *concrete_stored_files.File      `json:"metadata"`
	Blk *Block                           `json:"block"`
	Sig *concrete_stored_users.Signature `json:"signature"`
}

func createSignedBlock(met *concrete_stored_files.File, blk *Block, sig *concrete_stored_users.Signature) stored_validated_block.SignedBlock {
	out := SignedBlock{
		Met: met,
		Blk: blk,
		Sig: sig,
	}

	return &out
}

// GetMetaData returns the metadata
func (signedBlk *SignedBlock) GetMetaData() stored_files.File {
	return signedBlk.Met
}

// GetBlock returns the block
func (signedBlk *SignedBlock) GetBlock() stored_validated_block.Block {
	return signedBlk.Blk
}

// GetSignature returns the signature
func (signedBlk *SignedBlock) GetSignature() stored_users.Signature {
	return signedBlk.Sig
}
