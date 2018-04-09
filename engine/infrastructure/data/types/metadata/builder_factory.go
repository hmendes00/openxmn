package metadata

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
)

// BuilderFactory represents a concrete metadata builder factory implementation
type BuilderFactory struct {
}

// CreateBuilderFactory creates a new BuilderFactory instance
func CreateBuilderFactory() metadata.BuilderFactory {
	out := BuilderFactory{}
	return &out
}

// Create creates a new builder
func (fac *BuilderFactory) Create() metadata.Builder {
	out := createBuilder()
	return out
}
