package infrastructure

import (
	blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/domain"
	concrete_blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/infrastructure"
	validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
)

// Block represents a concrete validated Block implementation
type Block struct {
	Met *concrete_metadata.MetaData  `json:"metadata"`
	Blk *concrete_blocks.SignedBlock `json:"block"`
	LS  *concrete_users.Signatures   `json:"leader_signatures"`
}

func createBlock(met *concrete_metadata.MetaData, blk *concrete_blocks.SignedBlock, ls *concrete_users.Signatures) validated.Block {
	out := Block{
		Met: met,
		Blk: blk,
		LS:  ls,
	}

	return &out
}

// GetMetaData returns the MetaData
func (blk *Block) GetMetaData() metadata.MetaData {
	return blk.Met
}

// GetBlock returns the Block
func (blk *Block) GetBlock() blocks.SignedBlock {
	return blk.Blk
}

// GetSignatures returns the leader Signatures
func (blk *Block) GetSignatures() users.Signatures {
	return blk.LS
}
