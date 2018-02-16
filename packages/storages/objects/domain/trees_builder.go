package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// TreesBuilder represents a Trees builder
type TreesBuilder interface {
	Create() TreesBuilder
	WithTrees(trs []Tree) TreesBuilder
	WithHastTree(ht stored_files.File) TreesBuilder
	Now() (Trees, error)
}
