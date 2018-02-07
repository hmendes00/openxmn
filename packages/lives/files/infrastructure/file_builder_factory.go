package infrastructure

import (
	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
)

// FileBuilderFactory represents a concrete FileBuilderFactory implementation
type FileBuilderFactory struct {
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
