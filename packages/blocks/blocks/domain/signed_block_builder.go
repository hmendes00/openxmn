package domain

import (
	users "github.com/XMNBlockchain/core/packages/users/domain"
)

// SignedBlockBuilder represents a SignedBlock builder
type SignedBlockBuilder interface {
	Create() SignedBlockBuilder
	WithBlock(blk Block) SignedBlockBuilder
	WithSignature(sig users.Signature) SignedBlockBuilder
	Now() (SignedBlock, error)
}
