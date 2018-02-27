package domain

import (
	stored_validated_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/validated/domain"
)

// BlockService represents a Block service
type BlockService interface {
	Save(dirPath string, blk Block) (stored_validated_blocks.Block, error)
}
