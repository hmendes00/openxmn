package domain

import (
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// BlockService represents a Block service
type BlockService interface {
	Save(dirPath string, blk Block) (stored_objects.Tree, error)
}
