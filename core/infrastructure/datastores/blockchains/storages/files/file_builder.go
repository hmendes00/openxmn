package files

import (
	"errors"
	"time"

	dfil "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
)

type fileBuilder struct {
	path        string
	sizeInBytes int
	createdOn   *time.Time
}

func createFileBuilder() dfil.FileBuilder {
	out := fileBuilder{
		path:        "",
		sizeInBytes: 0,
		createdOn:   nil,
	}

	return &out
}

// Create initializes the file builder
func (build *fileBuilder) Create() dfil.FileBuilder {
	build.path = ""
	build.sizeInBytes = 0
	build.createdOn = nil
	return build
}

// WithPath adds a path to the file builder
func (build *fileBuilder) WithPath(path string) dfil.FileBuilder {
	build.path = path
	return build
}

// WithSizeInBytes adds a size in bytes to the file builder
func (build *fileBuilder) WithSizeInBytes(size int) dfil.FileBuilder {
	build.sizeInBytes = size
	return build
}

// CreatedOn adds a creation time to the file builder
func (build *fileBuilder) CreatedOn(ts time.Time) dfil.FileBuilder {
	build.createdOn = &ts
	return build
}

// Now builds a new File instance
func (build *fileBuilder) Now() (dfil.File, error) {
	if build.path == "" {
		return nil, errors.New("the path is mandatory in order to build a File instance")
	}

	if build.sizeInBytes == 0 {
		return nil, errors.New("the sizeInBytes is mandatory in order to build a File instance")
	}

	if build.createdOn == nil {
		return nil, errors.New("the createdOn is mandatory in order to build a File instance")
	}

	out := createFile(build.path, build.sizeInBytes, *build.createdOn)
	return out, nil
}
