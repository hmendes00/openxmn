package chunks

import (
	stored_chunks "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/chunks"
)

// Service represents a chunks service
type Service interface {
	Save(dirPath string, chk Chunks) (stored_chunks.Chunks, error)
}