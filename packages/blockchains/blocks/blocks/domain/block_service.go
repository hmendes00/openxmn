package domain

import (
	stored_block "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/domain"
)

// BlockService represents a block service
type BlockService interface {
	Save(dirPath string, blk Block) (stored_block.Block, error)
}
