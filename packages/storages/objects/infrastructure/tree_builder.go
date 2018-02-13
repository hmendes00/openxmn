package infrastructure

import (
	"errors"

	objs "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

type treeBuilder struct {
	name    string
	obj     objs.Object
	subObj  objs.Object
	subObjs []objs.Object
}

func createTreeBuilder() objs.TreeBuilder {
	out := treeBuilder{
		name:    "",
		obj:     nil,
		subObj:  nil,
		subObjs: nil,
	}
	return &out
}

// Create initializes the TreeBuilder
func (build *treeBuilder) Create() objs.TreeBuilder {
	build.name = ""
	build.obj = nil
	build.subObj = nil
	build.subObjs = nil
	return build
}

// WithName adds a name to the TreeBuilder
func (build *treeBuilder) WithName(name string) objs.TreeBuilder {
	build.name = name
	return build
}

// WithObject adds an object to the TreeBuilder
func (build *treeBuilder) WithObject(obj objs.Object) objs.TreeBuilder {
	build.obj = obj
	return build
}

// WithSubObject adds a sub object to the TreeBuilder
func (build *treeBuilder) WithSubObject(subObj objs.Object) objs.TreeBuilder {
	build.subObj = subObj
	return build
}

// WithSubObjects adds sub objects to the TreeBuilder
func (build *treeBuilder) WithSubObjects(subObjs []objs.Object) objs.TreeBuilder {
	build.subObjs = subObjs
	return build
}

// Now builds a new Tree instance
func (build *treeBuilder) Now() (objs.Tree, error) {
	if build.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Tree instance")
	}

	if build.obj != nil && build.subObj != nil && build.subObjs != nil {
		out := createTreeWithObjectWithSubObjectWithSubObjects(build.name, build.obj, build.subObj, build.subObjs)
		return out, nil
	}

	if build.obj != nil && build.subObj != nil {
		out := createTreeWithObjectWithSubObject(build.name, build.obj, build.subObj)
		return out, nil
	}

	if build.obj != nil && build.subObjs != nil {
		out := createTreeWithObjectWithSubObjects(build.name, build.obj, build.subObjs)
		return out, nil
	}

	if build.subObj != nil && build.subObjs != nil {
		out := createTreeWithSubObjectWithSubObjects(build.name, build.subObj, build.subObjs)
		return out, nil
	}

	if build.obj != nil {
		out := createTreeWithObject(build.name, build.obj)
		return out, nil
	}

	if build.subObj != nil {
		out := createTreeWithSubObject(build.name, build.subObj)
		return out, nil
	}

	if build.subObjs != nil {
		out := createTreeWithSubObjects(build.name, build.subObjs)
		return out, nil
	}

	return nil, errors.New("there must be at least an object, a sub-object or sub-objects.  None given")
}
