package domain

import (
	"time"

	files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// Chunks represents chunks stored on disk
type Chunks interface {
	GetHashTree() files.File
	GetChunks() []files.File
	CreatedOn() time.Time
}
