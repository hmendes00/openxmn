package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
)

// Command represents a concrete command implementation
type Command struct {
	Cmds *Commands `json:"commands"`
	Ins  *Insert   `json:"insert"`
	Up   *Update   `json:"update"`
	Del  *Delete   `json:"delete"`
	Err  *Error    `json:"error"`
}

func createCommandWithCommands(cmds *Commands) commands.Command {
	out := Command{
		Cmds: cmds,
		Ins:  nil,
		Up:   nil,
		Del:  nil,
		Err:  nil,
	}

	return &out
}

func createCommandWithInsert(ins *Insert) commands.Command {
	out := Command{
		Cmds: nil,
		Ins:  ins,
		Up:   nil,
		Del:  nil,
		Err:  nil,
	}

	return &out
}

func createCommandWithUpdate(up *Update) commands.Command {
	out := Command{
		Cmds: nil,
		Ins:  nil,
		Up:   up,
		Del:  nil,
		Err:  nil,
	}

	return &out
}

func createCommandWithDelete(del *Delete) commands.Command {
	out := Command{
		Cmds: nil,
		Ins:  nil,
		Up:   nil,
		Del:  del,
		Err:  nil,
	}

	return &out
}

func createCommandWithError(err *Error) commands.Command {
	out := Command{
		Cmds: nil,
		Ins:  nil,
		Up:   nil,
		Del:  nil,
		Err:  err,
	}

	return &out
}

// HasCommands returns true if the command is a Commands, false otherwise
func (cmd *Command) HasCommands() bool {
	return cmd.Cmds != nil
}

// GetCommands returns the Commands instance, if any
func (cmd *Command) GetCommands() commands.Commands {
	return cmd.Cmds
}

// HasInsert returns true if the command is an Insert, false otherwise
func (cmd *Command) HasInsert() bool {
	return cmd.Ins != nil
}

// GetInsert returns the Insert instance, if any
func (cmd *Command) GetInsert() commands.Insert {
	return cmd.Ins
}

// HasUpdate returns true if the command is an Update, false otherwise
func (cmd *Command) HasUpdate() bool {
	return cmd.Up != nil
}

// GetUpdate returns the Update instance, if any
func (cmd *Command) GetUpdate() commands.Update {
	return cmd.Up
}

// HasDelete returns true if the command is a Delete, false otherwise
func (cmd *Command) HasDelete() bool {
	return cmd.Del != nil
}

// GetDelete returns the Delete instance, if any
func (cmd *Command) GetDelete() commands.Delete {
	return cmd.Del
}

// HasError returns true if the command is an Error, false otherwise
func (cmd *Command) HasError() bool {
	return cmd.Err != nil
}

// GetError returns the Error instance, if any
func (cmd *Command) GetError() commands.Error {
	return cmd.Err
}
