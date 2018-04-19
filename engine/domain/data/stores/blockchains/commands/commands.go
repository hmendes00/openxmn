package commands

import (
	stored_chunks "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/chunks"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// Commands represents a stored commands
type Commands interface {
	GetMetaData() stored_files.File
	GetCommands() []stored_chunks.Chunks
}
