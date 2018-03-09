package domain

import (
	chained "github.com/XMNBlockchain/core/packages/blockchains/blocks/chained/domain"
	validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/domain"
)

// Chain represents the blockchain
type Chain interface {
	GetMetaData() MetaData
	GetFloorBlock() validated.Block
	HasCeilBlock() bool
	GetCeilBlock() chained.Block
}
