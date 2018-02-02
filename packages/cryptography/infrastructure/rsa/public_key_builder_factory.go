package rsa

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
)

// PublicKeyBuilderFactory represents a concrete PublicKeyBuilder factory
type PublicKeyBuilderFactory struct {
}

// CreatePublicKeyBuilderFactory creates a new PublicKeyBuilderFactory instance
func CreatePublicKeyBuilderFactory() cryptography.PublicKeyBuilderFactory {
	out := PublicKeyBuilderFactory{}
	return &out
}

// Create creates a new PublicKeyBuilder instance
func (fac *PublicKeyBuilderFactory) Create() cryptography.PublicKeyBuilder {
	out := createPublicKeyBuilder()
	return out
}
