package blocks

import (
	stored_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/users"
)

type signedBlock struct {
	metaData stored_files.File
	sig      stored_users.Signature
	blk      stored_blocks.Block
}

func createSignedBlock(metaData stored_files.File, sig stored_users.Signature, blk stored_blocks.Block) stored_blocks.SignedBlock {
	out := signedBlock{
		metaData: metaData,
		sig:      sig,
		blk:      blk,
	}

	return &out
}

// GetMetaData returns the metadata file
func (blk *signedBlock) GetMetaData() stored_files.File {
	return blk.metaData
}

// GetSignature returns the signature file
func (blk *signedBlock) GetSignature() stored_users.Signature {
	return blk.sig
}

// GetBlock returns the stored Block
func (blk *signedBlock) GetBlock() stored_blocks.Block {
	return blk.blk
}
