package chunks

import (
	files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/files"
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/hashtrees"
)

// Chunks represents a list of files to reproduce a total file
type Chunks interface {
	GetHashTree() hashtrees.HashTree
	GetChunks() []files.File
	GetData() ([]byte, error)
	Marshal(v interface{}) error
}