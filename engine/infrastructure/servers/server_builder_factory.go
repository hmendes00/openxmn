package servers

import (
	servers "github.com/XMNBlockchain/exmachina-network/engine/domain/servers"
)

// ServerBuilderFactory represents the concrete implementation of a ServerBuilderFactory
type ServerBuilderFactory struct {
}

// CreateServerBuilderFactory creates a new ServerBuilderFactory instance
func CreateServerBuilderFactory() servers.ServerBuilderFactory {
	out := ServerBuilderFactory{}
	return &out
}

// Create creates a new ServerBuilder instance
func (fac *ServerBuilderFactory) Create() servers.ServerBuilder {
	out := createServerBuilder()
	return out
}
