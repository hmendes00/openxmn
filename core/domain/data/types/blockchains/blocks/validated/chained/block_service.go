package domain

import (
	stored_chained_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/blocks/validated/chained"
)

// BlockService represents a block service
type BlockService interface {
	Save(dirPath string, blk Block) (stored_chained_blocks.Block, error)
}
