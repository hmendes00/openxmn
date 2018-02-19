package domain

import (
	"time"

	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

// SignedBlockBuilder represents a SignedBlock builder
type SignedBlockBuilder interface {
	Create() SignedBlockBuilder
	WithID(id *uuid.UUID) SignedBlockBuilder
	WithBlock(blk Block) SignedBlockBuilder
	WithSignature(sig users.Signature) SignedBlockBuilder
	CreatedOn(ts time.Time) SignedBlockBuilder
	Now() (SignedBlock, error)
}
