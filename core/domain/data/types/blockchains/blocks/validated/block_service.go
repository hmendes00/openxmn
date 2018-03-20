package domain

import (
	stored_validated_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/blocks/validated"
)

// BlockService represents a Block service
type BlockService interface {
	Save(dirPath string, blk Block) (stored_validated_blocks.Block, error)
}
