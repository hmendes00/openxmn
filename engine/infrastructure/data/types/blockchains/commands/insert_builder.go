package commands

import (
	"errors"

	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/commands/files"
)

type insertBuilder struct {
	fil files.File
}

func createInsertBuilder() commands.InsertBuilder {
	out := insertBuilder{
		fil: nil,
	}

	return &out
}

// Create initializes the InsertBuilder instance
func (build *insertBuilder) Create() commands.InsertBuilder {
	build.fil = nil
	return build
}

// WithFile adds a file to the InsertBuilder instance
func (build *insertBuilder) WithFile(fil files.File) commands.InsertBuilder {
	build.fil = fil
	return build
}

// Now builds a new Insert instance
func (build *insertBuilder) Now() (commands.Insert, error) {
	if build.fil == nil {
		return nil, errors.New("the file is mandatory in order to build an Insert instance")
	}

	out := createInsert(build.fil.(*concrete_files.File))
	return out, nil
}
