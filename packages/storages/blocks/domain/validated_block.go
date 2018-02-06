package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
)

// ValidatedBlock represents a stored block validated by peer users
type ValidatedBlock interface {
	GetHashTree() hashtrees.HashTree
	GetBlock() Block
	GetLeaderSignatures() []users.Signature
	CreatedOn() time.Time
}
