package infrastructure

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	user "github.com/XMNBlockchain/core/packages/lives/users/domain"
)

// SignatureBuilderFactory represents a concrete SignatureBuilderFactory
type SignatureBuilderFactory struct {
	sigBuilderFactory  cryptography.SignatureBuilderFactory
	userBuilderFactory user.UserBuilderFactory
}

// CreateSignatureBuilderFactory creates a new SignatureBuilderFactory instance
func CreateSignatureBuilderFactory(sigBuilderFactory cryptography.SignatureBuilderFactory, userBuilderFactory user.UserBuilderFactory) user.SignatureBuilderFactory {
	out := SignatureBuilderFactory{
		sigBuilderFactory:  sigBuilderFactory,
		userBuilderFactory: userBuilderFactory,
	}
	return &out
}

// Create creates a new SignatureBuilder instance
func (fac *SignatureBuilderFactory) Create() user.SignatureBuilder {
	out := createSignatureBuilder(fac.sigBuilderFactory, fac.userBuilderFactory)
	return out
}
