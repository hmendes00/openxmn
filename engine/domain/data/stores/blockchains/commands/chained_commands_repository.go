package commands

// ChainedCommandsRepository represents a stored chained commands repository
type ChainedCommandsRepository interface {
	Retrieve(dirPath string) (ChainedCommands, error)
}
