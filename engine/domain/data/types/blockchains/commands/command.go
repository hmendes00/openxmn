package commands

// Command represents a command
type Command interface {
	HasCommands() bool
	GetCommands() Commands
	HasInsert() bool
	GetInsert() Insert
	HasUpdate() bool
	GetUpdate() Update
	HasDelete() bool
	GetDelete() Delete
	HasError() bool
	GetError() Error
}
