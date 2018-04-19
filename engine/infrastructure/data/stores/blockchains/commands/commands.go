package commands

import (
	stored_commands "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/commands"
	stored_chunks "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/chunks"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	concrete_stored_chunks "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/chunks"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
)

// Commands represents a concrete stored Commands implementation
type Commands struct {
	Met  *concrete_stored_files.File      `json:"metadata"`
	Cmds []*concrete_stored_chunks.Chunks `json:"commands"`
}

func createCommands(met *concrete_stored_files.File, cmds []*concrete_stored_chunks.Chunks) stored_commands.Commands {
	out := Commands{
		Met:  met,
		Cmds: cmds,
	}

	return &out
}

// GetMetaData returns the metadata
func (cmd *Commands) GetMetaData() stored_files.File {
	return cmd.Met
}

// GetCommands returns the commands
func (cmd *Commands) GetCommands() []stored_chunks.Chunks {
	out := []stored_chunks.Chunks{}
	for _, oneStoredCmd := range cmd.Cmds {
		out = append(out, oneStoredCmd)
	}
	return out
}
