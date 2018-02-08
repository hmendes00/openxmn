package infrastructure

import (
	chunk "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
)

// ChunksBuilderFactory represents a concrete ChunksBuilderFactory implementation
type ChunksBuilderFactory struct {
}

// CreateChunksBuilderFactory builds a new ChunksBuilderFactory instance
func CreateChunksBuilderFactory() chunk.ChunksBuilderFactory {
	out := ChunksBuilderFactory{}
	return &out
}

// Create initializes the ChunksBuilder instance
func (fac *ChunksBuilderFactory) Create() chunk.ChunksBuilder {
	out := createChunksBuilder()
	return out
}
