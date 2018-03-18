package validated

import (
	stored_block "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks"
	stored_validated_block "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks/validated"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/users"
	concrete_stored_block "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/blocks"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/files"
	concrete_stored_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/users"
)

// Block represents a concrete stored validated block implementation
type Block struct {
	MetaData *concrete_stored_files.File        `json:"metadata"`
	Blk      *concrete_stored_block.SignedBlock `json:"signed_block"`
	Sigs     *concrete_stored_users.Signatures  `json:"signatures"`
}

func createBlock(metaData *concrete_stored_files.File, blk *concrete_stored_block.SignedBlock, sigs *concrete_stored_users.Signatures) stored_validated_block.Block {
	out := Block{
		MetaData: metaData,
		Blk:      blk,
		Sigs:     sigs,
	}

	return &out
}

// GetMetaData returns the metadata file
func (blk *Block) GetMetaData() stored_files.File {
	return blk.MetaData
}

// GetBlock returns the stored block
func (blk *Block) GetBlock() stored_block.SignedBlock {
	return blk.Blk
}

// GetSignatures returns the stored signatures
func (blk *Block) GetSignatures() stored_users.Signatures {
	return blk.Sigs
}