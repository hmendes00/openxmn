package commands

import (
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
)

// Delete represents a delete command
type Delete interface {
	GetFile() files.File
}
