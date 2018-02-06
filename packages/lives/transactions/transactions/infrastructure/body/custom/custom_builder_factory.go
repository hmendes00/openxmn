package custom

import (
	custom "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/custom"
)

// BuilderFactory represents a concrete CustomBuilderFactory
type BuilderFactory struct {
}

// CreateCustomBuilderFactory creates a new CustomBuilderFactory instance
func CreateCustomBuilderFactory() custom.BuilderFactory {
	out := BuilderFactory{}
	return &out
}

// Create creates a new CustomBuilder instance
func (fac *BuilderFactory) Create() custom.Builder {
	out := createCustomBuilder()
	return out
}
