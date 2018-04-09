package commands

import (
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
)

// Insert represents an insert command
type Insert interface {
	GetFile() files.File
}
