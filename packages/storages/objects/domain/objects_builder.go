package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// ObjectsBuilder represents an Objects builder
type ObjectsBuilder interface {
	Create() ObjectsBuilder
	WithHashTree(ht stored_files.File) ObjectsBuilder
	WithObjects(objs []Object) ObjectsBuilder
	Now() (Objects, error)
}
