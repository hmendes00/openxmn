package projects

import (
	projects "github.com/XMNBlockchain/exmachina-network/core/domain/projects"
)

// BuilderFactory represents a concrete project BuilderFactory implementation
type BuilderFactory struct {
}

// CreateBuilderFactory creates a new BuilderFactory instance
func CreateBuilderFactory() projects.BuilderFactory {
	out := BuilderFactory{}
	return &out
}

// Create creates a new projects Builder
func (fac *BuilderFactory) Create() projects.Builder {
	out := createBuilder()
	return out
}
