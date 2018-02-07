package infrastructure

import (
	"hash"
	"time"

	files "github.com/XMNBlockchain/core/packages/storages/files/domain"
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
func CreateFileBuilderFactory() files.FileBuilderFactory {
	out := FileBuilderFactory{}
	return &out
}

// Create creates a new FileBuilder instance
func (fac *FileBuilderFactory) Create() files.FileBuilder {
	out := createFileBuilder()
	return out
}