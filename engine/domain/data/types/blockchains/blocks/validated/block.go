package domain

import (
	blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// Block represents a block of transactions
type Block interface {
	GetMetaData() metadata.MetaData
	GetBlock() blocks.SignedBlock
	GetSignatures() users.Signatures
}
