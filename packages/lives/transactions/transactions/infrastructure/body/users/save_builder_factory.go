package users

import (
	users "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/users"
)

// SaveBuilderFactory represents a concrete SaveBuilderFactory
type SaveBuilderFactory struct {
}

// CreateSaveBuilderFactory creates a new SaveBuilderFactory instance
func CreateSaveBuilderFactory() users.SaveBuilderFactory {
	out := SaveBuilderFactory{}
	return &out
}

// Create creates a new SaveBuilder instance
func (fac *SaveBuilderFactory) Create() users.SaveBuilder {
	out := createSaveBuilder()
	return out
}
