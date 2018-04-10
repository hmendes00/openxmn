package tokens

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
)

// TokenBuilderFactory represents a concrete TokenBuilderFactory implementation
type TokenBuilderFactory struct {
	metaDataBuilderFactory metadata.BuilderFactory
}

// CreateTokenBuilderFactory creates a new TokenBuilderFactory instance
func CreateTokenBuilderFactory(metaDataBuilderFactory metadata.BuilderFactory) tokens.TokenBuilderFactory {
	out := TokenBuilderFactory{
		metaDataBuilderFactory: metaDataBuilderFactory,
	}
	return &out
}

// Create creates a new TokenBuilder instance
func (fac *TokenBuilderFactory) Create() tokens.TokenBuilder {
	out := createTokenBuilder(fac.metaDataBuilderFactory)
	return out
}
