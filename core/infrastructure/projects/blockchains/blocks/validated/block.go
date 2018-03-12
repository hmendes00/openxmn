package validated

import (
	blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/blocks"
	validated "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/blocks/validated"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/users"
	concrete_blocks "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/blocks"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/metadata"
	concrete_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/users"
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
