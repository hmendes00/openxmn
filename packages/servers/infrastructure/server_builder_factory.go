package infrastructure

import (
	servs "github.com/XMNBlockchain/core/packages/servers/domain"
)

// ServerBuilderFactory represents the concrete implementation of a ServerBuilderFactory
type ServerBuilderFactory struct {
}

// CreateServerBuilderFactory creates a new ServerBuilderFactory instance
func CreateServerBuilderFactory() servs.ServerBuilderFactory {
	out := ServerBuilderFactory{}
	return &out
}

// Create creates a new ServerBuilder instance
func (fac *ServerBuilderFactory) Create() servs.ServerBuilder {
	out := createServerBuilder()
	return out
}
