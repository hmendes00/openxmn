package infrastructure

import (
	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
)

// MetaDataBuilderFactory represents a concrete MetaDataBuilder instance
type MetaDataBuilderFactory struct {
}

// CreateMetaDataBuilderFactory creates a new MetaDataBuilderFactory instance
func CreateMetaDataBuilderFactory() objs.MetaDataBuilderFactory {
	out := MetaDataBuilderFactory{}
	return &out
}

// Create creates a MetaDataBuilder instance
func (fac *MetaDataBuilderFactory) Create() objs.MetaDataBuilder {
	out := createMetaDataBuilder()
	return out
}
