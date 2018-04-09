package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	uuid "github.com/satori/go.uuid"
)

// ChainedCommands represents a concrete chained commands implementation
type ChainedCommands struct {
	Met    *concrete_metadata.MetaData `json:"metadata"`
	Cmds   *Commands                   `json:"commands"`
	PrevID *uuid.UUID                  `json:"previous_id"`
	RootID *uuid.UUID                  `json:"root_id"`
}

func createChainedCommands(met *concrete_metadata.MetaData, cmds *Commands, prevID *uuid.UUID, rootID *uuid.UUID) commands.ChainedCommands {
	out := ChainedCommands{
		Met:    met,
		Cmds:   cmds,
		PrevID: prevID,
		RootID: rootID,
	}

	return &out
}

// GetMetaData returns the metadata
func (cmd *ChainedCommands) GetMetaData() metadata.MetaData {
	return cmd.Met
}

// GetCommands returns the commands
func (cmd *ChainedCommands) GetCommands() commands.Commands {
	return cmd.Cmds
}

// GetPreviousID returns the previous chained command ID
func (cmd *ChainedCommands) GetPreviousID() *uuid.UUID {
	return cmd.PrevID
}

// GetRootID returns the root command ID
func (cmd *ChainedCommands) GetRootID() *uuid.UUID {
	return cmd.RootID
}
