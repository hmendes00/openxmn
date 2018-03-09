package domain

import (
	validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/domain"
)

// Block represents a chained block
type Block interface {
	GetMetaData() MetaData
	GetBlock() validated.Block
}
