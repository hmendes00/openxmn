package commands

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// ChainedCommands represents a stored commands
type ChainedCommands interface {
	GetMetaData() stored_files.File
	GetCommands() Commands
	GetPreviousID() stored_files.File
	GetRootID() stored_files.File
}
