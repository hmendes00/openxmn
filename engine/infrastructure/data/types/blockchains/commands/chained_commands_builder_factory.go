package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
)

// ChainedCommandsBuilderFactory represents a concrete ChainedCommandsBuilderFactory instance
type ChainedCommandsBuilderFactory struct {
	metaDataBuilderFactory metadata.BuilderFactory
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
}

// CreateChainedCommandsBuilderFactory creates a new ChainedCommandsBuilderFactory instance
func CreateChainedCommandsBuilderFactory(metaDataBuilderFactory metadata.BuilderFactory, htBuilderFactory hashtrees.HashTreeBuilderFactory) commands.ChainedCommandsBuilderFactory {
	out := ChainedCommandsBuilderFactory{
		metaDataBuilderFactory: metaDataBuilderFactory,
		htBuilderFactory:       htBuilderFactory,
	}

	return &out
}

// Create creates a new ChainedCommandsBuilderFactory instance
func (fac *ChainedCommandsBuilderFactory) Create() commands.ChainedCommandsBuilder {
	out := createChainedCommandsBuilder(fac.metaDataBuilderFactory, fac.htBuilderFactory)
	return out
}
