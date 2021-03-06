package domain

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// MetaDataService represents a MetaData service
type MetaDataService interface {
	Save(dirPath string, met MetaData) (stored_files.File, error)
}
