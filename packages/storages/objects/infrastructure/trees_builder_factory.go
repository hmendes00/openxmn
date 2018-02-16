package infrastructure

import (
	objs "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// TreesBuilderFactory represents a concrete TreesBuilderFactory implementation
type TreesBuilderFactory struct {
}

// CreateTreesBuilderFactory creates a TreesBuilderFactory instance
func CreateTreesBuilderFactory() objs.TreesBuilderFactory {
	out := TreesBuilderFactory{}
	return &out
}

// Create creates a new Trees instance
func (fac *TreesBuilderFactory) Create() objs.TreesBuilder {
	out := createTreesBuilder()
	return out
}
