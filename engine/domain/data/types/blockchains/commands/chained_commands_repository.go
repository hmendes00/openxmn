package commands

// ChainedCommandsRepository represents a chained commands repository
type ChainedCommandsRepository interface {
	Retrieve(dirPath string) (ChainedCommands, error)
}
