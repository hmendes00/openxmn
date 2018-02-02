package rsa

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
)

// CreateSignatureBuilderFactoryForTests creates a SignatureBuilderFactory for tests
func CreateSignatureBuilderFactoryForTests() cryptography.SignatureBuilderFactory {
	publicKeyBuilderFactory := CreatePublicKeyBuilderFactory()
	out := CreateSignatureBuilderFactory(publicKeyBuilderFactory)
	return out
}
