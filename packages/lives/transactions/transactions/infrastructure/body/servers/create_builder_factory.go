package servers

import (
	servers "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/servers"
)

// CreateBuilderFactory represents a concrete CreateBuilderFactory
type CreateBuilderFactory struct {
}

// CreateCreateBuilderFactory creates a new CreateBuilderFactory instance
func CreateCreateBuilderFactory() servers.CreateBuilderFactory {
	out := CreateBuilderFactory{}
	return &out
}

// Create creates a new CreateBuilder instance
func (fac *CreateBuilderFactory) Create() servers.CreateBuilder {
	out := createCreateBuilder()
	return out
}
