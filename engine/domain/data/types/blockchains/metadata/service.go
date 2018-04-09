package metadata

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// Service represents an metadata service
type Service interface {
	Save(dirPath string, met MetaData) (stored_files.File, error)
}
