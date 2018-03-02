package domain

import (
	chained "github.com/XMNBlockchain/core/packages/blockchains/blocks/chained/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
)

// Chain represents the blockchain
type Chain interface {
	GetHashTree() hashtrees.HashTree
	GetMetaData() MetaData
	GetFloorBlock() chained.Block
	GetCeilBlock() chained.Block
}
