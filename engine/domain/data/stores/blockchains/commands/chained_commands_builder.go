package commands

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// ChainedCommandsBuilder represents a chained commands builder
type ChainedCommandsBuilder interface {
	Create() ChainedCommandsBuilder
	WithMetaData(met stored_files.File) ChainedCommandsBuilder
	WithCommands(cmds Commands) ChainedCommandsBuilder
	WithPreviousID(prevID stored_files.File) ChainedCommandsBuilder
	WithRootID(rootID stored_files.File) ChainedCommandsBuilder
	Now() (ChainedCommands, error)
}
