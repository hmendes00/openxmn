package commands

import (
	"errors"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	bills "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/bills"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
)

type builder struct {
	met      metadata.MetaData
	commands []commands.Command
	bi       bills.Bill
}

func createBuilder() commands.Builder {
	out := builder{
		met:      nil,
		commands: nil,
		bi:       nil,
	}

	return &out
}

// Create initializes the builder
func (build *builder) Create() commands.Builder {
	build.met = nil
	build.commands = nil
	build.bi = nil
	return build
}

// WithMetaData adds metadata to the builder
func (build *builder) WithMetaData(met metadata.MetaData) commands.Builder {
	build.met = met
	return build
}

// WithCommands adds commands to the builder
func (build *builder) WithCommands(cmds []commands.Command) commands.Builder {
	build.commands = cmds
	return build
}

// WithBill adds a bill to the builder
func (build *builder) WithBill(bi bills.Bill) commands.Builder {
	build.bi = bi
	return build
}

// Now builds a new Commands instance
func (build *builder) Now() (commands.Commands, error) {

	if build.met == nil {
		return nil, errors.New("the metadata is mandatory in order to build a Commands instance")
	}

	if build.commands == nil {
		return nil, errors.New("the []Command are mandatory in order to build a Commands instance")
	}

	cmds := []*Command{}
	for _, oneCmd := range build.commands {
		cmds = append(cmds, oneCmd.(*Command))
	}

	out := createCommands(build.met.(*concrete_metadata.MetaData), cmds)
	return out, nil
}
