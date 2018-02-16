package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// Trees represents a list of Tree
type Trees interface {
	GetTrees() []Tree
	GetHashTree() stored_files.File
}
