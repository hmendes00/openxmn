package domain

import (
	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
)

// ChunksService represents a chunks service
type ChunksService interface {
	Save(dirPath string, chk Chunks) (stored_chunks.Chunks, error)
	Delete(dirPath string, hash string) error
}
