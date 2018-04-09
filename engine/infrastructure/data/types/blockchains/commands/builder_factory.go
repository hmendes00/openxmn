package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
)

// BuilderFactory represents a concrete commands builder factory implementation
type BuilderFactory struct {
}

// CreateBuilderFactory creates a new BuilderFactory instance
func CreateBuilderFactory() commands.BuilderFactory {
	out := BuilderFactory{}
	return &out
}

// Create creates a new Builder instance
func (fac *BuilderFactory) Create() commands.Builder {
	out := createBuilder()
	return out
}
