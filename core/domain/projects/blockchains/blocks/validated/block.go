package domain

import (
	blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/blocks"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/users"
)

// Block represents a block of transactions
type Block interface {
	GetMetaData() metadata.MetaData
	GetBlock() blocks.SignedBlock
	GetSignatures() users.Signatures
}
