package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
)

// Commands represents a concrete commands implementation
type Commands struct {
	Met  *concrete_metadata.MetaData `json:"metadata"`
	Cmds []*Command                  `json:"commands"`
}

func createCommands(met *concrete_metadata.MetaData, cmds []*Command) commands.Commands {
	out := Commands{
		Met:  met,
		Cmds: cmds,
	}

	return &out
}

// GetMetaData returns the MetaData
func (cmds *Commands) GetMetaData() metadata.MetaData {
	return cmds.Met
}

// GetCommands returns the []Command
func (cmds *Commands) GetCommands() []commands.Command {
	out := []commands.Command{}
	for _, oneCmd := range cmds.Cmds {
		out = append(out, oneCmd)
	}

	return out
}
