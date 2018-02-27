package domain

import (
	stored_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/domain"
)

// SignedBlockService represents a signed block service
type SignedBlockService interface {
	Save(dirPath string, blk SignedBlock) (stored_blocks.SignedBlock, error)
}