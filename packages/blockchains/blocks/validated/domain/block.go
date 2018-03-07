package domain

import (
	blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
)

// Block represents a block of transactions
type Block interface {
	GetMetaData() metadata.MetaData
	GetBlock() blocks.SignedBlock
	GetSignatures() users.Signatures
}
