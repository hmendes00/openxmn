package servers

import (
	servers "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body/servers"
)

// DeleteBuilderFactory represents a concrete DeleteBuilderFactory
type DeleteBuilderFactory struct {
}

// CreateDeleteBuilderFactory creates a new DeleteBuilderFactory instance
func CreateDeleteBuilderFactory() servers.DeleteBuilderFactory {
	out := DeleteBuilderFactory{}
	return &out
}

// Create creates a new DeleteBuilder instance
func (fac *DeleteBuilderFactory) Create() servers.DeleteBuilder {
	out := createDeleteBuilder()
	return out
}
