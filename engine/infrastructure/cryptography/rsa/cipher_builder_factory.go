package rsa

import (
	cryptography "github.com/XMNBlockchain/exmachina-network/engine/domain/cryptography"
)

// CipherBuilderFactory represents a CipherBuilder factory
type CipherBuilderFactory struct {
	sigBuilderFactory cryptography.SignatureBuilderFactory
	pk                cryptography.PrivateKey
}

// CreateCipherBuilderFactory creates a new CipherBuilderFactory instance
func CreateCipherBuilderFactory(sigBuilderFactory cryptography.SignatureBuilderFactory, pk cryptography.PrivateKey) cryptography.CipherBuilderFactory {
	out := CipherBuilderFactory{
		sigBuilderFactory: sigBuilderFactory,
		pk:                pk,
	}

	return &out
}

// Create creates a new CipherBuilder instance
func (fac *CipherBuilderFactory) Create() cryptography.CipherBuilder {
	out := createCipherBuilder(fac.sigBuilderFactory, fac.pk)
	return out
}
