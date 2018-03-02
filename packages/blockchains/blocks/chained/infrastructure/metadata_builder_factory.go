package infrastructure

import (
	chained "github.com/XMNBlockchain/core/packages/blockchains/blocks/chained/domain"
)

// MetaDataBuilderFactory represents a concrete MetaDataBuilderFactory implementation
type MetaDataBuilderFactory struct {
}

// CreateMetaDataBuilderFactory creates a new MetaDataBuilderFactory instance
func CreateMetaDataBuilderFactory() chained.MetaDataBuilderFactory {
	out := MetaDataBuilderFactory{}
	return &out
}

// Create creates a new MetaDataBuilder instance
func (fac *MetaDataBuilderFactory) Create() chained.MetaDataBuilder {
	out := createMetaDataBuilder()
	return out
}
