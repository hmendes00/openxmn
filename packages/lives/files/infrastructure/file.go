package infrastructure

import (
	"hash"

	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
)

type file struct {
	h           hash.Hash
	sizeInBytes int
	data        []byte
	ext         string
}

func createFile(h hash.Hash, sizeInBytes int, data []byte, ext string) files.File {
	out := file{
		h:           h,
		sizeInBytes: sizeInBytes,
		data:        data,
		ext:         ext,
	}

	return &out
}

// GetHash returns the hash
func (fil *file) GetHash() hash.Hash {
	return fil.h
}

// GetSizeInBytes returns the size of the data in bytes
func (fil *file) GetSizeInBytes() int {
	return fil.sizeInBytes
}

// GetData returns the data
func (fil *file) GetData() []byte {
	return fil.data
}

// GetExtension returns the extension
func (fil *file) GetExtension() string {
	return fil.ext
}
