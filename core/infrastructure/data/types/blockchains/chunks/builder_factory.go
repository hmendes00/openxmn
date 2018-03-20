package chunks

import (
	chunk "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/chunks"
	files "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/files"
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/hashtrees"
)

// BuilderFactory represents a concrete BuilderFactory implementation
type BuilderFactory struct {
	fileBuilderFactory files.FileBuilderFactory
	htBuilderFactory   hashtrees.HashTreeBuilderFactory
	chkSizeInBytes     int
	extension          string
}

// CreateBuilderFactory creates a new BuilderFactory instance
func CreateBuilderFactory(fileBuilderFactory files.FileBuilderFactory, htBuilderFactory hashtrees.HashTreeBuilderFactory, chkSizeInBytes int, extension string) chunk.BuilderFactory {
	out := BuilderFactory{
		fileBuilderFactory: fileBuilderFactory,
		htBuilderFactory:   htBuilderFactory,
		chkSizeInBytes:     chkSizeInBytes,
		extension:          extension,
	}
	return &out
}

// Create creates a new Builder instance
func (fac *BuilderFactory) Create() chunk.Builder {
	out := createBuilder(fac.fileBuilderFactory, fac.htBuilderFactory, fac.chkSizeInBytes, fac.extension)
	return out
}
