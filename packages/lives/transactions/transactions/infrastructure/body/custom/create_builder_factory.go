package custom

import (
	custom "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/custom"
)

// CreateBuilderFactory represents a concrete CreateBuilderFactory
type CreateBuilderFactory struct {
}

// CreateCreateBuilderFactory creates a new CreateBuilderFactory instance
func CreateCreateBuilderFactory() custom.CreateBuilderFactory {
	out := CreateBuilderFactory{}
	return &out
}

// Create creates a new CreateBuilder instance
func (fac *CreateBuilderFactory) Create() custom.CreateBuilder {
	out := createCreateBuilder()
	return out
}
