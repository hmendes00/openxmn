package infrastructure

import (
	"errors"

	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	objs "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

type objectsBuilder struct {
	ht   stored_files.File
	objs []objs.Object
}

func createObjectsBuilder() objs.ObjectsBuilder {
	out := objectsBuilder{
		ht:   nil,
		objs: nil,
	}

	return &out
}

// Create initializes an ObjectsBuilder instance
func (build *objectsBuilder) Create() objs.ObjectsBuilder {
	build.ht = nil
	build.objs = nil
	return build
}

// WithHashTree adds an HashTree file to the ObjectsBuilder instance
func (build *objectsBuilder) WithHashTree(ht stored_files.File) objs.ObjectsBuilder {
	build.ht = ht
	return build
}

// WithObjects adds []Object instances to the ObjectsBuilder instance
func (build *objectsBuilder) WithObjects(objs []objs.Object) objs.ObjectsBuilder {
	build.objs = objs
	return build
}

// Now builds a new Objects instance
func (build *objectsBuilder) Now() (objs.Objects, error) {
	if build.ht == nil {
		return nil, errors.New("the hashtree file is mandatory in order to build an Objects instance")
	}

	if build.objs == nil {
		return nil, errors.New("the []Object are mandatory in order to build an Objects instance")
	}

	if len(build.objs) <= 0 {
		return nil, errors.New("the amount of []Object must be greater than 0")
	}

	out := createObjects(build.ht, build.objs)
	return out, nil
}
