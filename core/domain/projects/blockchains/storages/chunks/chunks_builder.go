package domain

import (
	"time"

	files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
)

// ChunksBuilder represents a Chunks builder
type ChunksBuilder interface {
	Create() ChunksBuilder
	WithHashTree(ht files.File) ChunksBuilder
	WithChunks(fil []files.File) ChunksBuilder
	CreatedOn(ts time.Time) ChunksBuilder
	Now() (Chunks, error)
}
