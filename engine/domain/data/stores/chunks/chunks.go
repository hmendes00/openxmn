package chunks

import (
	files "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/files"
)

// Chunks represents chunks stored on disk
type Chunks interface {
	GetHashTree() files.File
	GetChunks() []files.File
}
