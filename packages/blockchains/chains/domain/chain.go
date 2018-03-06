package domain

import (
	chained "github.com/XMNBlockchain/core/packages/blockchains/blocks/chained/domain"
)

// Chain represents the blockchain
type Chain interface {
	GetMetaData() MetaData
	GetFloorBlock() chained.Block
	GetCeilBlock() chained.Block
}
