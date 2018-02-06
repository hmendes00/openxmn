package users

import (
	users "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/users"
)

// DeleteBuilderFactory represents a concrete DeleteBuilderFactory
type DeleteBuilderFactory struct {
}

// CreateDeleteBuilderFactory creates a new DeleteBuilderFactory instance
func CreateDeleteBuilderFactory() users.DeleteBuilderFactory {
	out := DeleteBuilderFactory{}
	return &out
}

// Create creates a new DeleteBuilder instance
func (fac *DeleteBuilderFactory) Create() users.DeleteBuilder {
	out := createDeleteBuilder()
	return out
}
