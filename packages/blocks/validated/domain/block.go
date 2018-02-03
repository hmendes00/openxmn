package domain

import (
	blocks "github.com/XMNBlockchain/core/packages/blocks/blocks/domain"
	hashtree "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
)

// Block represents a block of transactions
type Block interface {
	GetHashTree() hashtree.Compact
	GetBlock() blocks.Block
	GetLeaderSignatures() []users.Signature
}
