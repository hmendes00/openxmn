package chained

import (
	stored_validated_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/blocks/validated"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/files"
)

// BlockBuilder represents a stored chained block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithMetaData(met stored_files.File) BlockBuilder
	WithBlock(blk stored_validated_blocks.Block) BlockBuilder
	Now() (Block, error)
}
