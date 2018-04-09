package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/commands/files"
)

// Delete represents a concrete Delete implementation
type Delete struct {
	Fil *concrete_files.File `json:"file"`
}

func createDelete(fil *concrete_files.File) commands.Delete {
	out := Delete{
		Fil: fil,
	}

	return &out
}

// GetFile returns the file
func (up *Delete) GetFile() files.File {
	return up.Fil
}
