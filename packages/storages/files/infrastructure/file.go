package infrastructure

import (
	"time"

	files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// File represents a concrete file representation
type File struct {
	Path        string    `json:"path"`
	Hash        string    `json:"hash"`
	SizeInBytes int       `json:"size"`
	CrOn        time.Time `json:"created_on"`
}

func createFile(path string, h string, sizeInBytes int, createdOn time.Time) files.File {
	out := File{
		Path:        path,
		Hash:        h,
		SizeInBytes: sizeInBytes,
		CrOn:        createdOn,
	}

	return &out
}

// GetPath returns the Path
func (fil *File) GetPath() string {
	return fil.Path
}

// GetHash returns the hash
func (fil *File) GetHash() string {
	return fil.Hash
}

// GetSizeInBytes returns the size in bytes
func (fil *File) GetSizeInBytes() int {
	return fil.SizeInBytes
}

// CreatedOn returns the creation time
func (fil *File) CreatedOn() time.Time {
	return fil.CrOn
}
