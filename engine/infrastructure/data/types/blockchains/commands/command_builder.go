package commands

import (
	"errors"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
)

type commandBuilder struct {
	met  metadata.MetaData
	cmds commands.Commands
	ins  commands.Insert
	up   commands.Update
	del  commands.Delete
	err  commands.Error
}

func createCommandBuilder() commands.CommandBuilder {
	out := commandBuilder{
		met:  nil,
		cmds: nil,
		ins:  nil,
		up:   nil,
		del:  nil,
		err:  nil,
	}

	return &out
}

// Create initializes the commandBuilder
func (build *commandBuilder) Create() commands.CommandBuilder {
	build.met = nil
	build.cmds = nil
	build.ins = nil
	build.up = nil
	build.del = nil
	build.err = nil
	return build
}

// WithMetaData adds a MetaData instance to the CommandBuilder instance
func (build *commandBuilder) WithMetaData(met metadata.MetaData) commands.CommandBuilder {
	build.met = met
	return build
}

// WithCommands adds a Commands instance to the CommandBuilder instance
func (build *commandBuilder) WithCommands(cmds commands.Commands) commands.CommandBuilder {
	build.cmds = cmds
	return build
}

// WithInsert adds an Insert instance to the CommandBuilder instance
func (build *commandBuilder) WithInsert(in commands.Insert) commands.CommandBuilder {
	build.ins = in
	return build
}

// WithUpdate adds an Update instance to the CommandBuilder instance
func (build *commandBuilder) WithUpdate(up commands.Update) commands.CommandBuilder {
	build.up = up
	return build
}

// WithDelete adds a Delete instance to the CommandBuilder instance
func (build *commandBuilder) WithDelete(del commands.Delete) commands.CommandBuilder {
	build.del = del
	return build
}

// Now builds a new Command instance
func (build *commandBuilder) Now() (commands.Command, error) {

	if build.met == nil {
		return nil, errors.New("the metadata is mandatory in order to build a Command instance")
	}

	if build.cmds != nil {
		out := createCommandWithCommands(build.met.(*concrete_metadata.MetaData), build.cmds.(*Commands))
		return out, nil
	}

	if build.ins != nil {
		out := createCommandWithInsert(build.met.(*concrete_metadata.MetaData), build.ins.(*Insert))
		return out, nil
	}

	if build.up != nil {
		out := createCommandWithUpdate(build.met.(*concrete_metadata.MetaData), build.up.(*Update))
		return out, nil
	}

	if build.del != nil {
		out := createCommandWithDelete(build.met.(*concrete_metadata.MetaData), build.del.(*Delete))
		return out, nil
	}

	return nil, errors.New("there must be a specific command (commands, insert, update, delete) in order to build a Command instance")
}
