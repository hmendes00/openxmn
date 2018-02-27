package domain

import (
	stored_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/domain"
)

// BlockService represents a Block service
type BlockService interface {
	Save(dirPath string, blk Block) (stored_blocks.Block, error)
}
