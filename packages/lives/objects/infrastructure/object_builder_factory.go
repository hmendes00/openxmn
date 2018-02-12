package infrastructure

import (
	objects "github.com/XMNBlockchain/core/packages/lives/objects/domain"
)

// ObjectBuilderFactory represents a concrete ObjectBuilderFactory implementation
type ObjectBuilderFactory struct {
}

// CreateObjectBuilderFactory creates a new ObjectBuilderFactory instance
func CreateObjectBuilderFactory() objects.ObjectBuilderFactory {
	out := ObjectBuilderFactory{}
	return &out
}

// Create creates a new ObjectBuilder instance
func (fac *ObjectBuilderFactory) Create() objects.ObjectBuilder {
	out := createObjectBuilder()
	return out
}
