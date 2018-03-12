package hashtrees

import (
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/hashtrees"
)

type hashTreeBuilderFactory struct {
}

// CreateHashTreeBuilderFactory creates a new hashTreeBuilderFactory instance
func CreateHashTreeBuilderFactory() hashtrees.HashTreeBuilderFactory {
	out := hashTreeBuilderFactory{}
	return &out
}

// Create creates a new hashTreeBuilder instance
func (fac *hashTreeBuilderFactory) Create() hashtrees.HashTreeBuilder {
	out := createHashTreeBuilder()
	return out
}
