package domain

import (
	"time"

	validated_blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
)

// BlockBuilder represents a chained BlockBuilder
type BlockBuilder interface {
	Create() BlockBuilder
	WithIndex(index int) BlockBuilder
	WithHashTree(ht hashtrees.HashTree) BlockBuilder
	WithBlock(blk validated_blocks.Block) BlockBuilder
	WithPreviousIndex(prevIndex int) BlockBuilder
	CreatedOn(ts time.Time) BlockBuilder
	Now() (Block, error)
}
