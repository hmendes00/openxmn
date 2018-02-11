package infrastructure

import (
	objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// ObjectBuilderFactory represents a concrete ObjectBuilderFactory imoplementation
type ObjectBuilderFactory struct {
}

func createObjectBuilderFactory() objects.ObjectBuilderFactory {
	out := ObjectBuilderFactory{}
	return &out
}

// Create creates a new ObjectBuilder instance
func (fac *ObjectBuilderFactory) Create() objects.ObjectBuilder {
	out := createObjectBuilder()
	return out
}
