package domain

import (
	validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
)

// Block represents a chained block
type Block interface {
	GetHashTree() hashtrees.HashTree
	GetMetaData() MetaData
	GetBlock() validated.Block
}
