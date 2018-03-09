package domain

import (
	stored_chained_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/chained/domain"
)

// BlockService represents a block service
type BlockService interface {
	Save(dirPath string, blk Block) (stored_chained_blocks.Block, error)
}
