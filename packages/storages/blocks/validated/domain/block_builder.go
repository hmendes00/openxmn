package domain

import (
	stored_block "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
)

// BlockBuilder represents a stored validated block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithMetaData(met stored_files.File) BlockBuilder
	WithBlock(blk stored_block.SignedBlock) BlockBuilder
	WithSignatures(sigs stored_users.Signatures) BlockBuilder
	Now() (Block, error)
}
