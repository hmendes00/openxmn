package domain

import (
	stored_validated_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks/validated"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
)

// BlockBuilder represents a stored chained block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithMetaData(met stored_files.File) BlockBuilder
	WithBlock(blk stored_validated_blocks.Block) BlockBuilder
	Now() (Block, error)
}
