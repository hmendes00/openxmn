package chunks

import (
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/files"
	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
)

// Chunks represents a list of files to reproduce a total file
type Chunks interface {
	GetHashTree() hashtrees.HashTree
	GetChunks() []files.File
	GetData() ([]byte, error)
	Marshal(v interface{}) error
}
