package domain

import (
	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	hashtree "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
)

// Block represents a block of transactions
type Block interface {
	GetHashTree() hashtree.HashTree
	GetBlock() blocks.SignedBlock
	GetLeaderSignatures() []users.Signature
}
