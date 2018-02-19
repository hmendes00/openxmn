package domain

import (
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// SignedBlockService represents a signed block service
type SignedBlockService interface {
	Save(dirPath string, blk SignedBlock) (stored_objects.Tree, error)
}
