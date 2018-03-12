package chunks

import (
	chunk "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/chunks"
)

// BuilderFactory represents a concrete ChunksBuilderFactory implementation
type BuilderFactory struct {
}

// CreateBuilderFactory builds a new BuilderFactory instance
func CreateBuilderFactory() chunk.ChunksBuilderFactory {
	out := BuilderFactory{}
	return &out
}

// Create initializes the ChunksBuilder instance
func (fac *BuilderFactory) Create() chunk.ChunksBuilder {
	out := createBuilder()
	return out
}
