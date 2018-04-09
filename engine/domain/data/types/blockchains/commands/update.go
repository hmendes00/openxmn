package commands

import (
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
)

// Update represents an update command
type Update interface {
	GetOriginalFile() files.File
	GetNewFile() files.File
}
