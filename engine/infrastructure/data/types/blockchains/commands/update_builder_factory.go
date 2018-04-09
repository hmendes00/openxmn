package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
)

// UpdateBuilderFactory represents a concrete UpdateBuilderFactory implementation
type UpdateBuilderFactory struct {
}

// CreateUpdateBuilderFactory creates a new UpdateBuilderFactory instance
func CreateUpdateBuilderFactory() commands.UpdateBuilderFactory {
	out := UpdateBuilderFactory{}
	return &out
}

// Create creates a new UpdateBuilder instance
func (fac *UpdateBuilderFactory) Create() commands.UpdateBuilder {
	out := createUpdateBuilder()
	return out
}
