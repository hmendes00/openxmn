package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/commands/files"
)

// Update represents a concrete Update implementation
type Update struct {
	OriginalFile *concrete_files.File `json:"original_file"`
	NewFile      *concrete_files.File `json:"new_file"`
}

func createUpdate(originalFile *concrete_files.File, newFile *concrete_files.File) commands.Update {
	out := Update{
		OriginalFile: originalFile,
		NewFile:      newFile,
	}

	return &out
}

// GetOriginalFile returns the original file
func (up *Update) GetOriginalFile() files.File {
	return up.OriginalFile
}

// GetNewFile returns the new file
func (up *Update) GetNewFile() files.File {
	return up.NewFile
}
