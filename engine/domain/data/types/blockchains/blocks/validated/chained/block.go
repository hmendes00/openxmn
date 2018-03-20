package domain

import (
	validated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
)

// Block represents a chained block
type Block interface {
	GetMetaData() MetaData
	GetBlock() validated.Block
}
