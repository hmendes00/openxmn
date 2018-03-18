package domain

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
)

// MetaDataService represents a MetaData service
type MetaDataService interface {
	Save(dirPath string, met MetaData) (stored_files.File, error)
}