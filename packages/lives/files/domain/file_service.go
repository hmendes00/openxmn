package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// FileService represents a file service
type FileService interface {
	Save(dirPath string, fil File) (stored_files.File, error)
	SaveAll(dirPath string, files []File) ([]stored_files.File, error)
	Delete(dirPath string, h string) error
	DeleteAll(dirPath string, h []string) error
}
