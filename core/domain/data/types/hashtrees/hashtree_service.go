package domain

import (
	stored_file "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/files"
)

// HashTreeService represents an HashTree service
type HashTreeService interface {
	Save(dirPath string, ht HashTree) (stored_file.File, error)
}
