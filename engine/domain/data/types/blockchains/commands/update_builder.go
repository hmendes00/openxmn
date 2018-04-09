package commands

import (
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
)

// UpdateBuilder represents an update builder
type UpdateBuilder interface {
	Create() UpdateBuilder
	WithOriginalFile(originalFile files.File) UpdateBuilder
	WithNewFile(newFile files.File) UpdateBuilder
	Now() (Update, error)
}
