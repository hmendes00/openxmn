package files

import (
	"time"

	dfil "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/files"
)

// File represents a concrete file representation
type File struct {
	Path        string    `json:"path"`
	SizeInBytes int       `json:"size"`
	CrOn        time.Time `json:"created_on"`
}

func createFile(path string, sizeInBytes int, createdOn time.Time) dfil.File {
	out := File{
		Path:        path,
		SizeInBytes: sizeInBytes,
		CrOn:        createdOn,
	}

	return &out
}

// GetPath returns the Path
func (fil *File) GetPath() string {
	return fil.Path
}

// GetSizeInBytes returns the size in bytes
func (fil *File) GetSizeInBytes() int {
	return fil.SizeInBytes
}

// CreatedOn returns the creation time
func (fil *File) CreatedOn() time.Time {
	return fil.CrOn
}
