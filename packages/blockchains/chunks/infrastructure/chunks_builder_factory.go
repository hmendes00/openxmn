package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	chunk "github.com/XMNBlockchain/core/packages/blockchains/chunks/domain"
	files "github.com/XMNBlockchain/core/packages/blockchains/files/domain"
)

// ChunksBuilderFactory represents a concrete ChunksBuilderFactory implementation
type ChunksBuilderFactory struct {
	fileBuilderFactory files.FileBuilderFactory
	htBuilderFactory   hashtrees.HashTreeBuilderFactory
	chkSizeInBytes     int
	extension          string
}

// CreateChunksBuilderFactory creates a new ChunksBuilderFactory instance
func CreateChunksBuilderFactory(fileBuilderFactory files.FileBuilderFactory, htBuilderFactory hashtrees.HashTreeBuilderFactory, chkSizeInBytes int, extension string) chunk.ChunksBuilderFactory {
	out := ChunksBuilderFactory{
		fileBuilderFactory: fileBuilderFactory,
		htBuilderFactory:   htBuilderFactory,
		chkSizeInBytes:     chkSizeInBytes,
		extension:          extension,
	}
	return &out
}

// Create creates a new ChunksBuilder instance
func (fac *ChunksBuilderFactory) Create() chunk.ChunksBuilder {
	out := createChunksBuilder(fac.fileBuilderFactory, fac.htBuilderFactory, fac.chkSizeInBytes, fac.extension)
	return out
}
