package commands

// CommandBuilder represents a command builder
type CommandBuilder interface {
	Create() CommandBuilder
	WithCommands(cmds Commands) CommandBuilder
	WithInsert(in Insert) CommandBuilder
	WithUpdate(up Update) CommandBuilder
	WithDelete(del Delete) CommandBuilder
	Now() (Command, error)
}
