package domain

import (
	stored_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks"
)

// SignedBlockService represents a signed block service
type SignedBlockService interface {
	Save(dirPath string, blk SignedBlock) (stored_blocks.SignedBlock, error)
}
