package commands

// ChainedCommandsBuilderFactory represents a chained commands builder factory
type ChainedCommandsBuilderFactory interface {
	Create() ChainedCommandsBuilder
}
