package domain

import (
	"time"

	files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
)

// Chunks represents chunks stored on disk
type Chunks interface {
	GetHashTree() files.File
	GetChunks() []files.File
	CreatedOn() time.Time
}
