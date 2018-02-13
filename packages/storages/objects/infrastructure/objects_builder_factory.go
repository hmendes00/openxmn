package infrastructure

import (
	objs "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// ObjectsBuilderFactory represents a concrete ObjectsBuilderFactory implementation
type ObjectsBuilderFactory struct {
}

// CreateObjectsBuilderFactory creates a new ObjectsBuilderFactory instance
func CreateObjectsBuilderFactory() objs.ObjectsBuilderFactory {
	out := ObjectsBuilderFactory{}
	return &out
}

// Create creates a new ObjectsBuilder instance
func (fac *ObjectsBuilderFactory) Create() objs.ObjectsBuilder {
	out := createObjectsBuilder()
	return out
}
