package domain

import (
	"time"

	validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	uuid "github.com/satori/go.uuid"
)

// BlockBuilder represents a chained block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithID(id *uuid.UUID) BlockBuilder
	WithIndex(index int) BlockBuilder
	WithPreviousIndex(prevIndex int) BlockBuilder
	WithHashTree(ht hashtrees.HashTree) BlockBuilder
	WithBlock(blk validated.Block) BlockBuilder
	CreatedOn(ts time.Time) BlockBuilder
	Now() (Block, error)
}
