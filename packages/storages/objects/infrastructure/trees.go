package infrastructure

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	objs "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

type trees struct {
	trs []objs.Tree
	ht  stored_files.File
}

func createTrees(trs []objs.Tree, ht stored_files.File) objs.Trees {
	out := trees{
		trs: trs,
		ht:  ht,
	}

	return &out
}

// GetTrees returns the []Tree
func (trs *trees) GetTrees() []objs.Tree {
	return trs.trs
}

// GetHashTree returns the HashTree
func (trs *trees) GetHashTree() stored_files.File {
	return trs.ht
}
