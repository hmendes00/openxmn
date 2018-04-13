package commands

import (
	"errors"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
)

type deleteBuilder struct {
	js []byte
}

func createDeleteBuilder() commands.DeleteBuilder {
	out := deleteBuilder{
		js: nil,
	}

	return &out
}

// Create initializes the DeleteBuilder instance
func (build *deleteBuilder) Create() commands.DeleteBuilder {
	build.js = nil
	return build
}

// WithFile adds json data to the DeleteBuilder instance
func (build *deleteBuilder) WithJS(js []byte) commands.DeleteBuilder {
	build.js = js
	return build
}

// Now builds a new Delete instance
func (build *deleteBuilder) Now() (commands.Delete, error) {
	if build.js == nil {
		return nil, errors.New("the json data is mandatory in order to build an Delete instance")
	}

	out := createDelete(build.js)
	return out, nil
}
