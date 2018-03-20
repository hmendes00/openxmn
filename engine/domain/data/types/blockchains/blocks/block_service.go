package domain

import (
	stored_block "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/blocks"
)

// BlockService represents a block service
type BlockService interface {
	Save(dirPath string, blk Block) (stored_block.Block, error)
}
