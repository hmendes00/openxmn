package files

import (
	"hash"
	"time"

	dfil "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
)

// FileBuilderFactory represents a concrete FileBuilderFactory implementation
type FileBuilderFactory struct {
	path        string
	h           hash.Hash
	sizeInBytes int
	contentType string
	createdOn   *time.Time
}

// CreateFileBuilderFactory creates a new FileBuilderFactory instance
func CreateFileBuilderFactory() dfil.FileBuilderFactory {
	out := FileBuilderFactory{}
	return &out
}

// Create creates a new FileBuilder instance
func (fac *FileBuilderFactory) Create() dfil.FileBuilder {
	out := createFileBuilder()
	return out
}
