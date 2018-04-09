package commands

import (
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
)

// InsertBuilder represents an insert builder
type InsertBuilder interface {
	Create() InsertBuilder
	WithFile(fil files.File) InsertBuilder
	Now() (Insert, error)
}
