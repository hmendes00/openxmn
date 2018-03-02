package infrastructure

import (
	stored_chains "github.com/XMNBlockchain/core/packages/storages/chains/domain"
)

// ChainBuilderFactory represents a concrete ChainBuilderFactory implementation
type ChainBuilderFactory struct {
}

// CreateChainBuilderFactory creates a new ChainBuilderFactory instance
func CreateChainBuilderFactory() stored_chains.ChainBuilderFactory {
	out := ChainBuilderFactory{}
	return &out
}

// Create creates a new ChainBuilder instance
func (fac *ChainBuilderFactory) Create() stored_chains.ChainBuilder {
	out := createChainBuilder()
	return out
}
