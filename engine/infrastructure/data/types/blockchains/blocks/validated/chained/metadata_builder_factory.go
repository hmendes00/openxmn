package chained

import (
	chained "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated/chained"
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
