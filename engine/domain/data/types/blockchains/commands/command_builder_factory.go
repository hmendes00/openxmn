package commands

// CommandBuilderFactory represents a command builder factory
type CommandBuilderFactory interface {
	Create() CommandBuilder
}
