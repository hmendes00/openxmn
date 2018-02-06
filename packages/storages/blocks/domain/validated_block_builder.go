package domain

import (
	"time"

	users "github.com/XMNBlockchain/core/packages/users/domain"
)

// ValidatedBlockBuilder represents a stored validated block builder
type ValidatedBlockBuilder interface {
	Create() ValidatedBlockBuilder
	WithBlock(blk Block) ValidatedBlockBuilder
	WithSignatures(sigs []users.Signature) ValidatedBlockBuilder
	CreatedOn(ts time.Time) ValidatedBlockBuilder
	Now() (ValidatedBlock, error)
}
