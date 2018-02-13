package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// Objects represents an Objects
type Objects interface {
	GetHashTree() stored_files.File
	GetObjects() []Object
}
