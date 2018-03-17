package chunks

import (
	files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
)

// Builder represents a Chunks builder
type Builder interface {
	Create() Builder
	WithHashTree(ht files.File) Builder
	WithChunks(fil []files.File) Builder
	Now() (Chunks, error)
}
