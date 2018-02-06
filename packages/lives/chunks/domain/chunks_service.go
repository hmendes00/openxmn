package domain

import (
	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
)

// ChunksService represents a chunks service
type ChunksService interface {
	Save(dirPath, chk Chunks) (stored_chunks.Chunks, error)
	SaveAll(dirPath, chks []Chunks) ([]stored_chunks.Chunks, error)
	Delete(dirPath string, h string) error
	DeleteAll(dirPath string, h []string) error
}
