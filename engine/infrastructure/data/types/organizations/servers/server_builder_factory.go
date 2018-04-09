package servers

import (
	organizations_server "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
)

// ServerBuilderFactory represents a concrete ServerBuilderFactory implementation
type ServerBuilderFactory struct {
}

// CreateServerBuilderFactory creates a new ServerBuilderFactory instance
func CreateServerBuilderFactory() organizations_server.ServerBuilderFactory {
	out := ServerBuilderFactory{}
	return &out
}

// Create builds a new ServerBuilder instance
func (fac *ServerBuilderFactory) Create() organizations_server.ServerBuilder {
	out := createServerBuilder()
	return out
}
