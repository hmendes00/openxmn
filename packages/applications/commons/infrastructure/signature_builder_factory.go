package infrastructure

import (
	commons "github.com/XMNBlockchain/core/packages/applications/commons/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
)

// SignatureBuilderFactory represents a concrete SignatureBuilder factory
type SignatureBuilderFactory struct {
	sigBuilderFactory users.SignatureBuilderFactory
}

// CreateSignatureBuilderFactory creates a new SignatureBuilderFactory instance
func CreateSignatureBuilderFactory(sigBuilderFactory users.SignatureBuilderFactory) commons.SignatureBuilderFactory {
	out := SignatureBuilderFactory{
		sigBuilderFactory: sigBuilderFactory,
	}

	return &out
}

// Create creates a new SignatureBuilder instance
func (fac *SignatureBuilderFactory) Create() commons.SignatureBuilder {
	out := createSignatureBuilder(fac.sigBuilderFactory)
	return out
}
