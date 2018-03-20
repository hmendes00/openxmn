package domain

import (
	blocks "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/blocks"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/metadata"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/users"
)

// Block represents a block of transactions
type Block interface {
	GetMetaData() metadata.MetaData
	GetBlock() blocks.SignedBlock
	GetSignatures() users.Signatures
}
