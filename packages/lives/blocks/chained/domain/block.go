package domain

import (
	"time"

	validated_blocks "github.com/XMNBlockchain/core/packages/lives/blocks/validated/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
)

// Block represents a chained block
type Block interface {
	GetIndex() int
	GetHashTree() hashtrees.HashTree
	GetBlock() validated_blocks.Block
	HasPrevious() bool
	GetPrevious() int
	CreatedOn() time.Time
}
