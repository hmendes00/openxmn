package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// MetaDataService represents an metadata service
type MetaDataService interface {
	Save(dirPath string, met MetaData) (stored_files.File, error)
}
