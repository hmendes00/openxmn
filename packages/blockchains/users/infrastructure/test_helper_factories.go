package infrastructure

import (
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	user "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
)

// CreateSignatureBuilderFactoryForTests creates a new SignatureBuilderFactory for tests
func CreateSignatureBuilderFactoryForTests() user.SignatureBuilderFactory {
	sigBuilderFactory := concrete_cryptography.CreateSignatureBuilderFactoryForTests()
	usrBuilderFactory := CreateUserBuilderFactory()
	out := CreateSignatureBuilderFactory(sigBuilderFactory, usrBuilderFactory)
	return out
}
