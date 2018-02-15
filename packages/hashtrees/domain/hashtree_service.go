package domain

import (
	stored_file "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// HashTreeService represents an HashTree service
type HashTreeService interface {
	Save(dirPath string, ht HashTree) (stored_file.File, error)
}
