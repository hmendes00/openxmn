package chunks

import (
	stored_chunks "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/chunks"
)

// Service represents a chunks service
type Service interface {
	Save(dirPath string, chk Chunks) (stored_chunks.Chunks, error)
}
