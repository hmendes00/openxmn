package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
)

type objects struct {
	ht   hashtrees.HashTree
	objs []objs.Object
}

func createObjects(ht hashtrees.HashTree, objs []objs.Object) objs.Objects {
	out := objects{
		ht:   ht,
		objs: objs,
	}

	return &out
}

// GetHashTree returns the HashTree
func (objs *objects) GetHashTree() hashtrees.HashTree {
	return objs.ht
}

// GetObjects returns the objects
func (objs *objects) GetObjects() []objs.Object {
	return objs.objs
}
