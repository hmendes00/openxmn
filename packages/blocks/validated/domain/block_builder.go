package domain

import (
	blocks "github.com/XMNBlockchain/core/packages/blocks/blocks/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
)

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithBlock(trs blocks.Block) BlockBuilder
	WithSignatures(sigs []users.Signature) BlockBuilder
	Now() (Block, error)
}
