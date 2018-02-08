package domain

import (
	"time"

	files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// ChunksBuilder represents a Chunks builder
type ChunksBuilder interface {
	Create() ChunksBuilder
	WithHashTree(ht files.File) ChunksBuilder
	WithChunks(fil []files.File) ChunksBuilder
	CreatedOn(ts time.Time) ChunksBuilder
	Now() (Chunks, error)
}
