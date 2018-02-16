package infrastructure

import (
	objs "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

type tree struct {
	name    string
	obj     objs.Object
	subObj  objs.Object
	subObjs objs.Objects
	subTrs  objs.Trees
}

func createTree(name string, obj objs.Object, subObj objs.Object, subObjs objs.Objects, subTrs objs.Trees) objs.Tree {
	out := tree{
		name:    name,
		obj:     obj,
		subObj:  subObj,
		subObjs: subObjs,
		subTrs:  subTrs,
	}

	return &out
}

// HasSubTrees returns true if there is a sub trees, false otherwise
func (tr *tree) HasSubTrees() bool {
	return tr.subTrs != nil
}

// GetSubTrees returns the sub Trees
func (tr *tree) GetSubTrees() objs.Trees {
	return tr.subTrs
}

// GetName returns the name
func (tr *tree) GetName() string {
	return tr.name
}

// HasObject returns true if there is an object, false otherwise
func (tr *tree) HasObject() bool {
	return tr.obj != nil
}

// GetObject returns the object, if any
func (tr *tree) GetObject() objs.Object {
	return tr.obj
}

// HasSubObject returns true if there is a sub object, false otherwise
func (tr *tree) HasSubObject() bool {
	return tr.subObj != nil
}

// GetSubObject returns the sub object, if any
func (tr *tree) GetSubObject() objs.Object {
	return tr.subObj
}

// HasSubObjects returns true if there is sub objects, false otherwise
func (tr *tree) HasSubObjects() bool {
	return tr.subObjs != nil
}

// GetSubObjects returns the sub objects, if any
func (tr *tree) GetSubObjects() objs.Objects {
	return tr.subObjs
}
