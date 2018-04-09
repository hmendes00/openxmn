package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	bills "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/bills"
)

type builder struct {
	commands []commands.Command
	bi       bills.Bill
}

func createBuilder() commands.Builder {
	out := builder{
		commands: nil,
		bi:       nil,
	}

	return &out
}

// Create initializes the builder
func (build *builder) Create() commands.Builder {
	build.commands = nil
	build.bi = nil
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

	return nil, nil
}
