package domain

import (
	stored_file "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// HashTreeService represents an HashTree service
type HashTreeService interface {
	Save(dirPath string, ht HashTree) (stored_file.File, error)
}
