package servers

import (
	servers "github.com/XMNBlockchain/core/packages/transactions/transactions/domain/body/servers"
)

// ServerBuilderFactory represents a concrete ServerBuilderFactory
type ServerBuilderFactory struct {
}

// CreateServerBuilderFactory creates a new ServerBuilderFactory instance
func CreateServerBuilderFactory() servers.BuilderFactory {
	out := ServerBuilderFactory{}
	return &out
}

// Create creates a new ServerBuilder instance
func (fac *ServerBuilderFactory) Create() servers.Builder {
	out := createServerBuilder()
	return out
}
