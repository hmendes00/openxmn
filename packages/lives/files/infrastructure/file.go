package infrastructure

import (
	"hash"

	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
)

// File represents a concrete file representation
type File struct {
	h           hash.Hash
	sizeInBytes int
	data        []byte
	ext         string
}

func createFile(h hash.Hash, sizeInBytes int, data []byte, ext string) files.File {
	out := File{
		h:           h,
		sizeInBytes: sizeInBytes,
		data:        data,
		ext:         ext,
	}

	return &out
}

// GetHash returns the hash
func (fil *File) GetHash() hash.Hash {
	return fil.h
}

// GetSizeInBytes returns the size of the data in bytes
func (fil *File) GetSizeInBytes() int {
	return fil.sizeInBytes
}

// GetData returns the data
func (fil *File) GetData() []byte {
	return fil.data
}

// GetExtension returns the extension
func (fil *File) GetExtension() string {
	return fil.ext
}
