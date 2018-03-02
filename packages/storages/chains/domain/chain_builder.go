package domain

import (
	stored_chained_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/chained/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// ChainBuilder represents a chain builder
type ChainBuilder interface {
	Create() ChainBuilder
	WithMetaData(met stored_files.File) ChainBuilder
	WithHashTree(ht stored_files.File) ChainBuilder
	WithFloorBlock(floorBlk stored_chained_blocks.Block) ChainBuilder
	WithCeilBlock(ceilBlk stored_chained_blocks.Block) ChainBuilder
	Now() (Chain, error)
}
