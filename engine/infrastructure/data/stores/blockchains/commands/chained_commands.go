package commands

import (
	stored_commands "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/commands"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
)

// ChainedCommands represents a concrete stored chained commands implementation
type ChainedCommands struct {
	Met    *concrete_stored_files.File `json:"metadata"`
	Cmds   *Commands                   `json:"commands"`
	PrevID *concrete_stored_files.File `json:"previous_id"`
	RootID *concrete_stored_files.File `json:"root_id"`
}

func createChainedCommands(
	met *concrete_stored_files.File,
	cmds *Commands,
	prevID *concrete_stored_files.File,
	rootID *concrete_stored_files.File,
) stored_commands.ChainedCommands {
	out := ChainedCommands{
		Met:    met,
		Cmds:   cmds,
		PrevID: prevID,
		RootID: rootID,
	}

	return &out
}

// GetMetaData returns the metadata
func (cmd *ChainedCommands) GetMetaData() stored_files.File {
	return cmd.Met
}

// GetCommands returns the commands
func (cmd *ChainedCommands) GetCommands() stored_commands.Commands {
	return cmd.Cmds
}

// GetPreviousID returns the previousID
func (cmd *ChainedCommands) GetPreviousID() stored_files.File {
	return cmd.PrevID
}

// GetRootID returns the rootID
func (cmd *ChainedCommands) GetRootID() stored_files.File {
	return cmd.RootID
}
