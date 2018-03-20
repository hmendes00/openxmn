package domain

import (
	blocks "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/blocks"
	metadata "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/users"
)

// Block represents a block of transactions
type Block interface {
	GetMetaData() metadata.MetaData
	GetBlock() blocks.SignedBlock
	GetSignatures() users.Signatures
}
