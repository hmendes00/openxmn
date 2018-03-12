package domain

import (
	validated "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/blocks/validated"
)

// Block represents a chained block
type Block interface {
	GetMetaData() MetaData
	GetBlock() validated.Block
}
