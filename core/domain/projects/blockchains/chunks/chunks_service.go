package domain

import (
	stored_chunks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/chunks"
)

// ChunksService represents a chunks service
type ChunksService interface {
	Save(dirPath string, chk Chunks) (stored_chunks.Chunks, error)
}
