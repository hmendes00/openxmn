package rsa

import (
	cryptography "github.com/XMNBlockchain/exmachina-network/core/domain/cryptography"
)

// SignatureBuilderFactory represents a concrete SignatureBuilder factory
type SignatureBuilderFactory struct {
	publicKeyBuilderFactory cryptography.PublicKeyBuilderFactory
}

// CreateSignatureBuilderFactory creates a new SignatureBuilderFactory instance
func CreateSignatureBuilderFactory(publicKeyBuilderFactory cryptography.PublicKeyBuilderFactory) cryptography.SignatureBuilderFactory {
	out := SignatureBuilderFactory{
		publicKeyBuilderFactory: publicKeyBuilderFactory,
	}

	return &out
}

// Create creates a new SignatureBuilder instance
func (fac *SignatureBuilderFactory) Create() cryptography.SignatureBuilder {
	out := createSignatureBuilder(fac.publicKeyBuilderFactory)
	return out
}
