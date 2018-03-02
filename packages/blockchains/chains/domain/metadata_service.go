package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// MetaDataService represents the metadata service
type MetaDataService interface {
	Save(met MetaData) (stored_files.File, error)
}
