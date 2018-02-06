package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
)

// ChainedBlockBuilder represents a chained block builder
type ChainedBlockBuilder interface {
	Create() ChainedBlockBuilder
	WithIndex(index int) ChainedBlockBuilder
	WithHashTree(ht hashtrees.HashTree) ChainedBlockBuilder
	WithBlock(blk ValidatedBlock) ChainedBlockBuilder
	WithPreviousBlock(prevBlk ValidatedBlock) ChainedBlockBuilder
	CreatedOn(ts time.Time) ChainedBlockBuilder
	Now() (ChainedBlock, error)
}
