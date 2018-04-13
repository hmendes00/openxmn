package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
)

// Delete represents a concrete Delete implementation
type Delete struct {
	JS []byte `json:"js"`
}

func createDelete(js []byte) commands.Delete {
	out := Delete{
		JS: js,
	}

	return &out
}

// GetJS returns the json data
func (up *Delete) GetJS() []byte {
	return up.JS
}
