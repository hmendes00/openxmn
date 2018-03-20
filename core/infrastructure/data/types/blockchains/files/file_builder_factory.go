package files

import (
	files "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/files"
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
