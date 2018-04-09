package files

import (
	"errors"
	"hash"

	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/files"
	uuid "github.com/satori/go.uuid"
)

type fileBuilder struct {
	id   *uuid.UUID
	typ  string
	hash string
}

func createFileBuilder() files.FileBuilder {
	out := fileBuilder{
		id:   nil,
		typ:  "",
		hash: "",
	}

	return &out
}

// Create initializes the FileBuilder instance
func (build *fileBuilder) Create() files.FileBuilder {
	build.id = nil
	build.typ = ""
	build.hash = ""
	return build
}

// WithID adds an ID to the FileBuilder
func (build *fileBuilder) WithID(id *uuid.UUID) files.FileBuilder {
	build.id = id
	return build
}

// WithType adds a type to the FileBuilder
func (build *fileBuilder) WithType(typ string) files.FileBuilder {
	build.typ = typ
	return build
}

// WithHash adds an hash to the FileBuilder
func (build *fileBuilder) WithHash(h hash.Hash) files.FileBuilder {
	build.hash = string(h.Sum(nil))
	return build
}

// Now builds a new File instance
func (build *fileBuilder) Now() (files.File, error) {
	if build.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a File instance")
	}

	if build.typ == "" {
		return nil, errors.New("the type is mandatory in order to build a File instance")
	}

	if build.hash == "" {
		return nil, errors.New("the hash is mandatory in order to build a File instance")
	}

	out := createFile(build.id, build.typ, build.hash)
	return out, nil
}
