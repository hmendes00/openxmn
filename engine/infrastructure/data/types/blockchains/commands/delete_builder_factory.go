package commands

import (
	commands "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands"
)

// DeleteBuilderFactory represents a concrete delete builder factory implementation
type DeleteBuilderFactory struct {
}

// CreateDeleteBuilderFactory creates a new DeleteBuilderFactory instance
func CreateDeleteBuilderFactory() commands.DeleteBuilderFactory {
	out := DeleteBuilderFactory{}
	return &out
}

// Create creates a new DeleteBuilder instance
func (fac *DeleteBuilderFactory) Create() commands.DeleteBuilder {
	out := createDeleteBuilder()
	return out
}
