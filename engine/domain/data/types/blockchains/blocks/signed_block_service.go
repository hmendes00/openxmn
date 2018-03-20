package domain

import (
	stored_blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/blocks"
)

// SignedBlockService represents a signed block service
type SignedBlockService interface {
	Save(dirPath string, blk SignedBlock) (stored_blocks.SignedBlock, error)
}
