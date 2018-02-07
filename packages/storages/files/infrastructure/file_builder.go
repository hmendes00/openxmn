package infrastructure

import (
	"errors"
	"hash"
	"time"

	files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

type fileBuilder struct {
	path        string
	h           hash.Hash
	sizeInBytes int
	createdOn   *time.Time
}

func createFileBuilder() files.FileBuilder {
	out := fileBuilder{
		path:        "",
		h:           nil,
		sizeInBytes: 0,
		createdOn:   nil,
	}

	return &out
}

// Create initializes the file builder
func (build *fileBuilder) Create() files.FileBuilder {
	build.path = ""
	build.h = nil
	build.sizeInBytes = 0
	build.createdOn = nil
	return build
}

// WithPath adds a path to the file builder
func (build *fileBuilder) WithPath(path string) files.FileBuilder {
	build.path = path
	return build
}

// WithHash adds an hash to the file builder
func (build *fileBuilder) WithHash(h hash.Hash) files.FileBuilder {
	build.h = h
	return build
}

// WithSizeInBytes adds a size in bytes to the file builder
func (build *fileBuilder) WithSizeInBytes(size int) files.FileBuilder {
	build.sizeInBytes = size
	return build
}

// CreatedOn adds a creation time to the file builder
func (build *fileBuilder) CreatedOn(ts time.Time) files.FileBuilder {
	build.createdOn = &ts
	return build
}

// Now builds a new File instance
func (build *fileBuilder) Now() (files.File, error) {
	if build.path == "" {
		return nil, errors.New("the path is mandatory in order to build a File instance")
	}

	if build.h == nil {
		return nil, errors.New("the hash is mandatory in order to build a File instance")
	}

	if build.sizeInBytes == 0 {
		return nil, errors.New("the sizeInBytes is mandatory in order to build a File instance")
	}

	if build.createdOn == nil {
		return nil, errors.New("the createdOn is mandatory in order to build a File instance")
	}

	out := createFile(build.path, build.h, build.sizeInBytes, *build.createdOn)
	return out, nil
}
