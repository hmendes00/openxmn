package domain

import (
	"time"

	validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	uuid "github.com/satori/go.uuid"
)

// Block represents a chained block
type Block interface {
	GetID() *uuid.UUID
	GetIndex() int
	GetPreviousIndex() int
	GetHashTree() hashtrees.HashTree
	GetBlock() validated.Block
	CreatedOn() time.Time
}
