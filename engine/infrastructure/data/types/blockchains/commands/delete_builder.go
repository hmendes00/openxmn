package commands

import (
	"errors"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/commands/files"
)

type deleteBuilder struct {
	fil files.File
}

func createDeleteBuilder() commands.DeleteBuilder {
	out := deleteBuilder{
		fil: nil,
	}

	return &out
}

// Create initializes the DeleteBuilder instance
func (build *deleteBuilder) Create() commands.DeleteBuilder {
	build.fil = nil
	return build
}

// WithFile adds a file to the DeleteBuilder instance
func (build *deleteBuilder) WithFile(fil files.File) commands.DeleteBuilder {
	build.fil = fil
	return build
}

// Now builds a new Delete instance
func (build *deleteBuilder) Now() (commands.Delete, error) {
	if build.fil == nil {
		return nil, errors.New("the file is mandatory in order to build an Delete instance")
	}

	out := createDelete(build.fil.(*concrete_files.File))
	return out, nil
}
