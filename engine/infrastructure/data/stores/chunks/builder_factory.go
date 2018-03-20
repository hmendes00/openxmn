package chunks

import (
	chunk "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/chunks"
)

// BuilderFactory represents a concrete ChunksBuilderFactory implementation
type BuilderFactory struct {
}

// CreateBuilderFactory builds a new BuilderFactory instance
func CreateBuilderFactory() chunk.BuilderFactory {
	out := BuilderFactory{}
	return &out
}

// Create initializes the ChunksBuilder instance
func (fac *BuilderFactory) Create() chunk.Builder {
	out := createBuilder()
	return out
}
