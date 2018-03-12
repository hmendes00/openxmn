package metadata

import (
	met "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
)

// MetaDataBuilderFactory represents a concrete MetaDataBuilder instance
type MetaDataBuilderFactory struct {
}

// CreateMetaDataBuilderFactory creates a new MetaDataBuilderFactory instance
func CreateMetaDataBuilderFactory() met.MetaDataBuilderFactory {
	out := MetaDataBuilderFactory{}
	return &out
}

// Create creates a MetaDataBuilder instance
func (fac *MetaDataBuilderFactory) Create() met.MetaDataBuilder {
	out := createMetaDataBuilder()
	return out
}
