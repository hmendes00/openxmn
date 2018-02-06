package domain

import (
	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
)

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithBlock(blk blocks.SignedBlock) BlockBuilder
	WithSignatures(sigs []users.Signature) BlockBuilder
	Now() (Block, error)
}
