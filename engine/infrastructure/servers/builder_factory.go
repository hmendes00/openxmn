package servers

import (
	servers "github.com/XMNBlockchain/exmachina-network/engine/domain/servers"
)

// BuilderFactory represents a concrete BuilderFactory implementation
type BuilderFactory struct {
}

// CreateBuilderFactory creates a new CreateBuilderFactory instance
func CreateBuilderFactory() servers.BuilderFactory {
	out := BuilderFactory{}
	return &out
}

// Create creates a new Servers instance
func (fac *BuilderFactory) Create() servers.Builder {
	out := createBuilder()
	return out
}
