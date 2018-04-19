package commands

import (
	"errors"

	stored_commands "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/commands"
	stored_chunks "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/chunks"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	concrete_stored_chunks "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/chunks"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
)

type builder struct {
	met  stored_files.File
	cmds []stored_chunks.Chunks
}

func createBuilder() stored_commands.Builder {
	out := builder{
		met:  nil,
		cmds: nil,
	}

	return &out
}

// Create initializes the commands builder
func (build *builder) Create() stored_commands.Builder {
	build.met = nil
	build.cmds = nil
	return build
}

// WithMetaData adds metadata to the commands builder
func (build *builder) WithMetaData(met stored_files.File) stored_commands.Builder {
	build.met = met
	return build
}

// WithCommands adds a commands to the commands builder
func (build *builder) WithCommands(cmds []stored_chunks.Chunks) stored_commands.Builder {
	build.cmds = cmds
	return build
}

// Now builds a new commands instance
func (build *builder) Now() (stored_commands.Commands, error) {
	if build.met == nil {
		return nil, errors.New("the metadata is mandatory in order to build a commands instance")
	}

	if build.cmds == nil {
		return nil, errors.New("the []Command is mandatory in order to build a commands instance")
	}

	cmds := []*concrete_stored_chunks.Chunks{}
	for _, oneCmd := range build.cmds {
		cmds = append(cmds, oneCmd.(*concrete_stored_chunks.Chunks))
	}

	out := createCommands(build.met.(*concrete_stored_files.File), cmds)
	return out, nil
}
