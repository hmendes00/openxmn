package commands

import (
	stored_commands "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/commands"
)

// BuilderFactory represents a concrete commands BuilderFactory implementation
type BuilderFactory struct {
}

// CreateBuilderFactory creates a new commands BuilderFactory instance
func CreateBuilderFactory() stored_commands.BuilderFactory {
	out := BuilderFactory{}
	return &out
}

// Create creates a new builder instance
func (fac *BuilderFactory) Create() stored_commands.Builder {
	out := createBuilder()
	return out
}
