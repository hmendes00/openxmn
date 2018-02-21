package infrastructure

import (
	met "github.com/XMNBlockchain/core/packages/lives/metadata/domain"
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
