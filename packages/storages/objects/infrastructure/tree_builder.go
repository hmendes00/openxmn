package infrastructure

import (
	"errors"

	objs "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

type treeBuilder struct {
	name    string
	obj     objs.Object
	subObj  objs.Object
	subObjs objs.Objects
	subTrs  objs.Trees
}

func createTreeBuilder() objs.TreeBuilder {
	out := treeBuilder{
		name:    "",
		obj:     nil,
		subObj:  nil,
		subObjs: nil,
		subTrs:  nil,
	}
	return &out
}

// Create initializes the TreeBuilder
func (build *treeBuilder) Create() objs.TreeBuilder {
	build.name = ""
	build.obj = nil
	build.subObj = nil
	build.subObjs = nil
	build.subTrs = nil
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
func (build *treeBuilder) WithSubObjects(subObjs objs.Objects) objs.TreeBuilder {
	build.subObjs = subObjs
	return build
}

// WithTrees adds sub objects to the TreeBuilder
func (build *treeBuilder) WithSubTrees(trs objs.Trees) objs.TreeBuilder {
	build.subTrs = trs
	return build
}

// Now builds a new Tree instance
func (build *treeBuilder) Now() (objs.Tree, error) {
	if build.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Tree instance")
	}

	out := createTree(build.name, build.obj, build.subObj, build.subObjs, build.subTrs)
	return out, nil
}
