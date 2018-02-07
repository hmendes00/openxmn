package infrastructure

import (
	"hash"
	"time"

	files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

type file struct {
	path        string
	h           hash.Hash
	sizeInBytes int
	createdOn   time.Time
}

func createFile(path string, h hash.Hash, sizeInBytes int, createdOn time.Time) files.File {
	out := file{
		path:        path,
		h:           h,
		sizeInBytes: sizeInBytes,
		createdOn:   createdOn,
	}

	return &out
}

// GetPath returns the Path
func (fil *file) GetPath() string {
	return fil.path
}

// GetHash returns the hash
func (fil *file) GetHash() hash.Hash {
	return fil.h
}

// GetSizeInBytes returns the size in bytes
func (fil *file) GetSizeInBytes() int {
	return fil.sizeInBytes
}

// CreatedOn returns the creation time
func (fil *file) CreatedOn() time.Time {
	return fil.createdOn
}
