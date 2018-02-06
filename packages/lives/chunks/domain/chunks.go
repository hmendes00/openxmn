package domain

import (
	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// Chunks represents a list of files to reproduce a total file
type Chunks interface {
	GetHashTree() hashtrees.HashTree
	GetChunks() []files.File
}
