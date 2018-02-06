package users

import (
	users "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/users"
)

// CreateBuilderFactory represents a concrete CreateBuilderFactory
type CreateBuilderFactory struct {
}

// CreateCreateBuilderFactory creates a new CreateBuilderFactory instance
func CreateCreateBuilderFactory() users.CreateBuilderFactory {
	out := CreateBuilderFactory{}
	return &out
}

// Create creates a new CreateBuilder instance
func (fac *CreateBuilderFactory) Create() users.CreateBuilder {
	out := createCreateBuilder()
	return out
}
