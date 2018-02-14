package infrastructure

import (
	objs "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

type tree struct {
	name    string
	obj     objs.Object
	subObj  objs.Object
	subObjs objs.Objects
}

func createTreeWithObject(name string, obj objs.Object) objs.Tree {
	out := tree{
		name:    name,
		obj:     obj,
		subObj:  nil,
		subObjs: nil,
	}

	return &out
}

func createTreeWithSubObject(name string, subObj objs.Object) objs.Tree {
	out := tree{
		name:    name,
		obj:     nil,
		subObj:  subObj,
		subObjs: nil,
	}

	return &out
}

func createTreeWithSubObjects(name string, subObjs objs.Objects) objs.Tree {
	out := tree{
		name:    name,
		obj:     nil,
		subObj:  nil,
		subObjs: subObjs,
	}

	return &out
}

func createTreeWithObjectWithSubObject(name string, obj objs.Object, subObj objs.Object) objs.Tree {
	out := tree{
		name:    name,
		obj:     obj,
		subObj:  subObj,
		subObjs: nil,
	}

	return &out
}

func createTreeWithObjectWithSubObjects(name string, obj objs.Object, subObjs objs.Objects) objs.Tree {
	out := tree{
		name:    name,
		obj:     obj,
		subObj:  nil,
		subObjs: subObjs,
	}

	return &out
}

func createTreeWithSubObjectWithSubObjects(name string, subObj objs.Object, subObjs objs.Objects) objs.Tree {
	out := tree{
		name:    name,
		obj:     nil,
		subObj:  subObj,
		subObjs: subObjs,
	}

	return &out
}

func createTreeWithObjectWithSubObjectWithSubObjects(name string, obj objs.Object, subObj objs.Object, subObjs objs.Objects) objs.Tree {
	out := tree{
		name:    name,
		obj:     obj,
		subObj:  subObj,
		subObjs: subObjs,
	}

	return &out
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
