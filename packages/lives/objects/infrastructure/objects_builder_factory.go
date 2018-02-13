package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
)

// ObjectsBuilderFactory represents a concrete ObjectsBuilderFactory implementation
type ObjectsBuilderFactory struct {
	htBuilderFactory hashtrees.HashTreeBuilderFactory
}

// CreateObjectsBuilderFactory creates a new ObjectsBuilderFactory instance
func CreateObjectsBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory) objs.ObjectsBuilderFactory {
	out := ObjectsBuilderFactory{
		htBuilderFactory: htBuilderFactory,
	}
	return &out
}

// Create creates a new ObjectsBuilder instance
func (fac *ObjectsBuilderFactory) Create() objs.ObjectsBuilder {
	out := createObjectsBuilder(fac.htBuilderFactory)
	return out
}
