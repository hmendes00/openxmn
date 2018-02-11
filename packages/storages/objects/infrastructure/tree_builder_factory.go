package infrastructure

import (
	objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// TreeBuilderFactory represents a concrete TreeBuilderFactory imoplementation
type TreeBuilderFactory struct {
}

func createTreeBuilderFactory() objects.TreeBuilderFactory {
	out := TreeBuilderFactory{}
	return &out
}

// Create creates a new TreeBuilder instance
func (fac *TreeBuilderFactory) Create() objects.TreeBuilder {
	out := createTreeBuilder()
	return out
}
