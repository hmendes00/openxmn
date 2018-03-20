package domain

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
)

// FileService represents a file service
type FileService interface {
	Save(dirPath string, fil File) (stored_files.File, error)
	SaveAll(dirPath string, files []File) ([]stored_files.File, error)
	Delete(dirPath string, fileName string) error
	DeleteAll(dirPath string) error
}
