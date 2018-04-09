package domain

import (
	stored_validated_blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/blocks/validated"
)

// SignedBlockService represents a SignedBlock service
type SignedBlockService interface {
	Save(dirPath string, signedBlk SignedBlock) (stored_validated_blocks.SignedBlock, error)
}
