package domain

import (
	stored_validated_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/validated/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// BlockBuilder represents a stored chained block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithMetaData(met stored_files.File) BlockBuilder
	WithHashTree(ht stored_files.File) BlockBuilder
	WithBlock(blk stored_validated_blocks.Block) BlockBuilder
	Now() (Block, error)
}
