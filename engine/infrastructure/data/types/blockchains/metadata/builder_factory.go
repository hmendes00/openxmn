package metadata

import (
	met "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
)

// BuilderFactory represents a concrete Builder instance
type BuilderFactory struct {
}

// CreateBuilderFactory creates a new BuilderFactory instance
func CreateBuilderFactory() met.BuilderFactory {
	out := BuilderFactory{}
	return &out
}

// Create creates a Builder instance
func (fac *BuilderFactory) Create() met.Builder {
	out := createBuilder()
	return out
}
