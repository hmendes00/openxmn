package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
)

// Update represents a concrete Update implementation
type Update struct {
	OriginalJS []byte `json:"original_js"`
	NewJS      []byte `json:"new_js"`
}

func createUpdate(originalJS []byte, newJS []byte) commands.Update {
	out := Update{
		OriginalJS: originalJS,
		NewJS:      newJS,
	}

	return &out
}

// GetOriginalJS returns the original json data
func (up *Update) GetOriginalJS() []byte {
	return up.OriginalJS
}

// GetNewJS returns the new json data
func (up *Update) GetNewJS() []byte {
	return up.NewJS
}
