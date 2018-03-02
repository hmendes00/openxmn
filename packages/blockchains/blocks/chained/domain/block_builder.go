package domain

import (
	validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/domain"
)

// BlockBuilder represents a chained block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithMetaData(met MetaData) BlockBuilder
	WithBlock(blk validated.Block) BlockBuilder
	Now() (Block, error)
}
