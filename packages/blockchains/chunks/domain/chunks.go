package domain

import (
	files "github.com/XMNBlockchain/core/packages/blockchains/files/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
)

// Chunks represents a list of files to reproduce a total file
type Chunks interface {
	GetHashTree() hashtrees.HashTree
	GetChunks() []files.File
	Marshal(v interface{}) error
}
