package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
)

// ChainedBlock represents a chained block
type ChainedBlock interface {
	GetIndex() int
	GetHashTree() hashtrees.HashTree
	GetBlock() Block
	HasPrevious() bool
	GetPrevious() int
	CreatedOn() time.Time
}
