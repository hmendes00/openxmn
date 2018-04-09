package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/commands/files"
)

// Insert represents a concrete Insert implementation
type Insert struct {
	Fil *concrete_files.File `json:"file"`
}

func createInsert(fil *concrete_files.File) commands.Insert {
	out := Insert{
		Fil: fil,
	}

	return &out
}

// GetFile returns the file
func (up *Insert) GetFile() files.File {
	return up.Fil
}
