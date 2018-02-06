package body

import (
	body "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body"
)

// BuilderFactory represents a concrete BodyBuilderFactory implementation
type BuilderFactory struct {
}

// CreateBodyBuilderFactory creates a new BodyBuilderFactory instance
func CreateBodyBuilderFactory() body.BuilderFactory {
	out := BuilderFactory{}
	return &out
}

// Create creates a new BodyBuilder instance
func (fac *BuilderFactory) Create() body.Builder {
	build := createBodyBuilder()
	return build
}
