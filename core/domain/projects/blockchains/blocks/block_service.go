package domain

import (
	stored_block "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks"
)

// BlockService represents a block service
type BlockService interface {
	Save(dirPath string, blk Block) (stored_block.Block, error)
}
