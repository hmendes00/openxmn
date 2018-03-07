package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	user "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
)

// SignaturesBuilderFactory represents a concrete SignaturesBuilderFactory implementation
type SignaturesBuilderFactory struct {
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
}

// CreateSignaturesBuilderFactory creates a new SignaturesBuilderFactory instance
func CreateSignaturesBuilderFactory(htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) user.SignaturesBuilderFactory {
	out := SignaturesBuilderFactory{
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
	}

	return &out
}

// Create builds a new SignaturesBuilder instance
func (fac *SignaturesBuilderFactory) Create() user.SignaturesBuilder {
	out := createSignaturesBuilder(fac.htBuilderFactory, fac.metaDataBuilderFactory)
	return out
}
