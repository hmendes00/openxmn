package commands

import (
	stored_commands "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/commands"
)

// ChainedCommandsBuilderFactory represents a concrete ChainedCommandsBuilderFactory implementation
type ChainedCommandsBuilderFactory struct {
}

// CreateChainedCommandsBuilderFactory creates a new ChainedCommandsBuilderFactory instance
func CreateChainedCommandsBuilderFactory() stored_commands.ChainedCommandsBuilderFactory {
	out := ChainedCommandsBuilderFactory{}
	return &out
}

// Create creates a new ChainedCommandsBuilder instance
func (fac *ChainedCommandsBuilderFactory) Create() stored_commands.ChainedCommandsBuilder {
	out := createChainedCommandsBuilder()
	return out
}
