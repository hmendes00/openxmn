package blocks

import (
	stored_blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/blocks"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/users"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
	concrete_stored_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/users"
)

// SignedBlock represents a concrete stored signed block implementation
type SignedBlock struct {
	MetaData *concrete_stored_files.File      `json:"metadata"`
	Sig      *concrete_stored_users.Signature `json:"signature"`
	Blk      *Block                           `json:"block"`
}

func createSignedBlock(metaData *concrete_stored_files.File, sig *concrete_stored_users.Signature, blk *Block) stored_blocks.SignedBlock {
	out := SignedBlock{
		MetaData: metaData,
		Sig:      sig,
		Blk:      blk,
	}

	return &out
}

// GetMetaData returns the metadata file
func (blk *SignedBlock) GetMetaData() stored_files.File {
	return blk.MetaData
}

// GetSignature returns the signature file
func (blk *SignedBlock) GetSignature() stored_users.Signature {
	return blk.Sig
}

// GetBlock returns the stored Block
func (blk *SignedBlock) GetBlock() stored_blocks.Block {
	return blk.Blk
}
