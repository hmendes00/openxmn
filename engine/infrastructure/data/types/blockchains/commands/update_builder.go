package commands

import (
	"errors"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
)

type updateBuilder struct {
	originalJS []byte
	newJS      []byte
}

func createUpdateBuilder() commands.UpdateBuilder {
	out := updateBuilder{
		originalJS: nil,
		newJS:      nil,
	}

	return &out
}

// Create initializes the UpdateBuilder instance
func (build *updateBuilder) Create() commands.UpdateBuilder {
	build.originalJS = nil
	build.newJS = nil
	return build
}

// WithOriginalJS adds json data to the UpdateBuilder instance
func (build *updateBuilder) WithOriginalJS(originalJS []byte) commands.UpdateBuilder {
	build.originalJS = originalJS
	return build
}

// WithNewJS adds new json data to the UpdateBuilder instance
func (build *updateBuilder) WithNewJS(newJS []byte) commands.UpdateBuilder {
	build.newJS = newJS
	return build
}

// Now builds a new Update instance
func (build *updateBuilder) Now() (commands.Update, error) {
	if build.originalJS == nil {
		return nil, errors.New("the original json data is mandatory in order to build an Update instance")
	}

	if build.newJS == nil {
		return nil, errors.New("the new json data is mandatory in order to build an Update instance")
	}

	out := createUpdate(build.originalJS, build.newJS)
	return out, nil
}
