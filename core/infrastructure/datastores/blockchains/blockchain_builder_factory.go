package blockchains

import (
	blockchains "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains"
)

// BlockchainBuilderFactory represents a concrete BlockchainBuilderFactory implementation
type BlockchainBuilderFactory struct {
}

// CreateBlockchainBuilderFactory creates a new BlockchainBuilderFactory instance
func CreateBlockchainBuilderFactory() blockchains.BlockchainBuilderFactory {
	out := BlockchainBuilderFactory{}
	return &out
}

// Create creates a new blockchain builder
func (fac *BlockchainBuilderFactory) Create() blockchains.BlockchainBuilder {
	out := createBlockchainBuilder()
	return out
}
