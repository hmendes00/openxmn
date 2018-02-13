package infrastructure

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	objs "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

type objects struct {
	ht   stored_files.File
	objs []objs.Object
}

func createObjects(ht stored_files.File, objs []objs.Object) objs.Objects {
	out := objects{
		ht:   ht,
		objs: objs,
	}

	return &out
}

// GetHashTree returns the hashtree file
func (obj *objects) GetHashTree() stored_files.File {
	return obj.ht
}

// GetObjects returns the []Object
func (obj *objects) GetObjects() []objs.Object {
	return obj.objs
}
