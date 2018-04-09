package users

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	user "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// SignaturesBuilderFactory represents a concrete SignaturesBuilderFactory implementation
type SignaturesBuilderFactory struct {
	metaDataBuilderFactory metadata.BuilderFactory
}

// CreateSignaturesBuilderFactory creates a new SignaturesBuilderFactory instance
func CreateSignaturesBuilderFactory(metaDataBuilderFactory metadata.BuilderFactory) user.SignaturesBuilderFactory {
	out := SignaturesBuilderFactory{
		metaDataBuilderFactory: metaDataBuilderFactory,
	}

	return &out
}

// Create builds a new SignaturesBuilder instance
func (fac *SignaturesBuilderFactory) Create() user.SignaturesBuilder {
	out := createSignaturesBuilder(fac.metaDataBuilderFactory)
	return out
}
