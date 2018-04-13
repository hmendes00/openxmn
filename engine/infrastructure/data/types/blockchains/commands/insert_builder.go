package commands

import (
	"errors"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
)

type insertBuilder struct {
	js []byte
}

func createInsertBuilder() commands.InsertBuilder {
	out := insertBuilder{
		js: nil,
	}

	return &out
}

// Create initializes the InsertBuilder instance
func (build *insertBuilder) Create() commands.InsertBuilder {
	build.js = nil
	return build
}

// WithJS adds json data to the InsertBuilder instance
func (build *insertBuilder) WithJS(js []byte) commands.InsertBuilder {
	build.js = js
	return build
}

// Now builds a new Insert instance
func (build *insertBuilder) Now() (commands.Insert, error) {
	if build.js == nil {
		return nil, errors.New("the json data is mandatory in order to build an Insert instance")
	}

	out := createInsert(build.js)
	return out, nil
}
