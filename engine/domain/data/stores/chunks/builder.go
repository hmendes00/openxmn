package chunks

import (
	files "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/files"
)

// Builder represents a Chunks builder
type Builder interface {
	Create() Builder
	WithHashTree(ht files.File) Builder
	WithChunks(fil []files.File) Builder
	Now() (Chunks, error)
}
