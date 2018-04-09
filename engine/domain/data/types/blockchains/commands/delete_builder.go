package commands

import (
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
)

// DeleteBuilder represents a delete builder
type DeleteBuilder interface {
	Create() DeleteBuilder
	WithFile(fil files.File) DeleteBuilder
	Now() (Delete, error)
}
