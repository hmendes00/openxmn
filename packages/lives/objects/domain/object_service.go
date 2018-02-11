package domain

import (
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// ObjectService represents an object service
type ObjectService interface {
	Save(dirPath string, obj Object) (stored_objects.Object, error)
}
