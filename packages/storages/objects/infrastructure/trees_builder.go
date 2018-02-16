package infrastructure

import (
	"errors"

	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	objs "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

type treesBuilder struct {
	trs []objs.Tree
	ht  stored_files.File
}

func createTreesBuilder() objs.TreesBuilder {
	out := treesBuilder{
		trs: nil,
		ht:  nil,
	}

	return &out
}

// Create initializes the Trees builder
func (build *treesBuilder) Create() objs.TreesBuilder {
	build.trs = nil
	build.ht = nil
	return build
}

// WithTrees adds a []Tree to the TreesBuilder
func (build *treesBuilder) WithTrees(trs []objs.Tree) objs.TreesBuilder {
	build.trs = trs
	return build
}

// WithHastTree adds an HashTree to the TreesBuilder
func (build *treesBuilder) WithHastTree(ht stored_files.File) objs.TreesBuilder {
	build.ht = ht
	return build
}

// Now builds a new Trees instance
func (build *treesBuilder) Now() (objs.Trees, error) {
	if build.trs == nil {
		return nil, errors.New("the []Tree is mandatory in order to build a Trees instance")
	}

	if len(build.trs) <= 0 {
		return nil, errors.New("the []Tree must contain at least 1 Tree instance")
	}

	if build.ht == nil {
		return nil, errors.New("the HashTree is mandatory in order to build a Trees instance")
	}

	out := createTrees(build.trs, build.ht)
	return out, nil
}
