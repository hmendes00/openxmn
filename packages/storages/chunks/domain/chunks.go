package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// Chunks represents chunks stored on disk
type Chunks interface {
	GetHashTree() hashtrees.HashTree
	GetChunks() []files.File
	CreatedOn() time.Time
}
