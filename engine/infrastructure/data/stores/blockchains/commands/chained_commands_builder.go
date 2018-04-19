package commands

import (
	"errors"

	stored_commands "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/commands"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
)

type chainedCommandsBuilder struct {
	met    stored_files.File
	cmds   stored_commands.Commands
	prevID stored_files.File
	rootID stored_files.File
}

func createChainedCommandsBuilder() stored_commands.ChainedCommandsBuilder {
	out := chainedCommandsBuilder{
		met:    nil,
		cmds:   nil,
		prevID: nil,
		rootID: nil,
	}

	return &out
}

// Create initializes a ChainedCommandsBuilder instance
func (build *chainedCommandsBuilder) Create() stored_commands.ChainedCommandsBuilder {
	build.met = nil
	build.cmds = nil
	build.prevID = nil
	build.rootID = nil
	return build
}

// WithMetaData adds metadata to the chained commands builder
func (build *chainedCommandsBuilder) WithMetaData(met stored_files.File) stored_commands.ChainedCommandsBuilder {
	build.met = met
	return build
}

// WithCommands adds a commands to the chained commands builder
func (build *chainedCommandsBuilder) WithCommands(cmds stored_commands.Commands) stored_commands.ChainedCommandsBuilder {
	build.cmds = cmds
	return build
}

// WithPreviousID adds a previousID to the chained commands builder
func (build *chainedCommandsBuilder) WithPreviousID(prevID stored_files.File) stored_commands.ChainedCommandsBuilder {
	build.prevID = prevID
	return build
}

// WithRootID adds a rootID to the chained commands builder
func (build *chainedCommandsBuilder) WithRootID(rootID stored_files.File) stored_commands.ChainedCommandsBuilder {
	build.rootID = rootID
	return build
}

// Now builds a new ChainedCommands instance
func (build *chainedCommandsBuilder) Now() (stored_commands.ChainedCommands, error) {
	if build.met == nil {
		return nil, errors.New("the metadata is mandatory in order to build a ChainedCommands instance")
	}

	if build.cmds == nil {
		return nil, errors.New("the commands is mandatory in order to build a ChainedCommands instance")
	}

	if build.prevID == nil {
		return nil, errors.New("the previousID is mandatory in order to build a ChainedCommands instance")
	}

	if build.rootID == nil {
		return nil, errors.New("the rootID is mandatory in order to build a ChainedCommands instance")
	}

	out := createChainedCommands(build.met.(*concrete_stored_files.File), build.cmds.(*Commands), build.prevID.(*concrete_stored_files.File), build.rootID.(*concrete_stored_files.File))
	return out, nil
}
