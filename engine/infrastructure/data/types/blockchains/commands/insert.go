package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
)

// Insert represents a concrete Insert implementation
type Insert struct {
	JS []byte `json:"js"`
}

func createInsert(js []byte) commands.Insert {
	out := Insert{
		JS: js,
	}

	return &out
}

// GetJS returns the json data
func (up *Insert) GetJS() []byte {
	return up.JS
}
