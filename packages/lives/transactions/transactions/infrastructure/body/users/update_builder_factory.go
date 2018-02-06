package users

import "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/users"

// UpdateBuilderFactory represents a concrete UpdateBuilderFactory
type UpdateBuilderFactory struct {
}

// CreateUpdateBuilderFactory creates a new UpdateBuilderFactory instance
func CreateUpdateBuilderFactory() users.UpdateBuilderFactory {
	out := UpdateBuilderFactory{}
	return &out
}

// Create creates a new CreateBuilder instance
func (fac *UpdateBuilderFactory) Create() users.UpdateBuilder {
	out := createUpdateBuilder()
	return out
}
