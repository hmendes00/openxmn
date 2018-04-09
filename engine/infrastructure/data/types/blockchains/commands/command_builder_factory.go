package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
)

// CommandBuilderFactory creates a new CommandBuilderFactory instance
type CommandBuilderFactory struct {
}

// CreateCommandBuilderFactory creates a new CommandBuilderFactory instance
func CreateCommandBuilderFactory() commands.CommandBuilderFactory {
	out := CommandBuilderFactory{}
	return &out
}

// Create creates a new CommandBuilderFactory instance
func (fac *CommandBuilderFactory) Create() commands.CommandBuilder {
	out := createCommandBuilder()
	return out
}
