package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
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
