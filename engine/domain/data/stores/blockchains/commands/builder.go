package commands

import (
	stored_chunks "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/chunks"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// Builder represents a stored commands builder
type Builder interface {
	Create() Builder
	WithMetaData(met stored_files.File) Builder
	WithCommands(cmds []stored_chunks.Chunks) Builder
	Now() (Commands, error)
}
