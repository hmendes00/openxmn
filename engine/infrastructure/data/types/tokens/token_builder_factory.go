package tokens

import (
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
)

// TokenBuilderFactory represents a concrete TokenBuilderFactory implementation
type TokenBuilderFactory struct {
}

// CreateTokenBuilderFactory creates a new TokenBuilderFactory instance
func CreateTokenBuilderFactory() tokens.TokenBuilderFactory {
	out := TokenBuilderFactory{}
	return &out
}

// Create creates a new TokenBuilder instance
func (fac *TokenBuilderFactory) Create() tokens.TokenBuilder {
	out := createTokenBuilder()
	return out
}
