package infrastructure

import (
	objs "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// TreeBuilderFactory represents a concrete TreeBuilderFactory imoplementation
type TreeBuilderFactory struct {
}

// CreateTreeBuilderFactory creates a TreeBuilderFactory instance
func CreateTreeBuilderFactory() objs.TreeBuilderFactory {
	out := TreeBuilderFactory{}
	return &out
}

// Create creates a new TreeBuilder instance
func (fac *TreeBuilderFactory) Create() objs.TreeBuilder {
	out := createTreeBuilder()
	return out
}
