package domain

import (
	"time"

	validated_blocks "github.com/XMNBlockchain/core/packages/lives/blocks/validated/domain"
)

// BlockBuilder represents a chained BlockBuilder
type BlockBuilder interface {
	Create() BlockBuilder
	WithIndex(index int) BlockBuilder
	WithBlock(blk validated_blocks.Block) BlockBuilder
	WithPreviousIndex(prevIndex int) BlockBuilder
	CreatedOn(ts time.Time) BlockBuilder
	Now() (Block, error)
}
