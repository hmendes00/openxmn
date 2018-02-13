package domain

import (
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// ObjectsService represents an objects service
type ObjectsService interface {
	Save(dirPath string, obj Objects) (stored_objects.Objects, error)
}
