package commands

import (
	"errors"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/commands/files"
)

type updateBuilder struct {
	originalFile files.File
	newFile      files.File
}

func createUpdateBuilder() commands.UpdateBuilder {
	out := updateBuilder{
		originalFile: nil,
		newFile:      nil,
	}

	return &out
}

// Create initializes the UpdateBuilder instance
func (build *updateBuilder) Create() commands.UpdateBuilder {
	build.originalFile = nil
	build.newFile = nil
	return build
}

// WithOriginalFile adds an original file to the UpdateBuilder instance
func (build *updateBuilder) WithOriginalFile(originalFile files.File) commands.UpdateBuilder {
	build.originalFile = originalFile
	return build
}

// WithNewFile adds a new file to the UpdateBuilder instance
func (build *updateBuilder) WithNewFile(newFile files.File) commands.UpdateBuilder {
	build.newFile = newFile
	return build
}

// Now builds a new Update instance
func (build *updateBuilder) Now() (commands.Update, error) {
	if build.originalFile == nil {
		return nil, errors.New("the originalFile is mandatory in order to build an Update instance")
	}

	if build.newFile == nil {
		return nil, errors.New("the newFile is mandatory in order to build an Update instance")
	}

	out := createUpdate(build.originalFile.(*concrete_files.File), build.newFile.(*concrete_files.File))
	return out, nil
}
