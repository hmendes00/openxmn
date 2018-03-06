package infrastructure

import (
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	user "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
)

// SignatureBuilderFactory represents a concrete SignatureBuilderFactory
type SignatureBuilderFactory struct {
	sigBuilderFactory      cryptography.SignatureBuilderFactory
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
}

// CreateSignatureBuilderFactory creates a new SignatureBuilderFactory instance
func CreateSignatureBuilderFactory(sigBuilderFactory cryptography.SignatureBuilderFactory, htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) user.SignatureBuilderFactory {
	out := SignatureBuilderFactory{
		sigBuilderFactory:      sigBuilderFactory,
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
	}
	return &out
}

// Create creates a new SignatureBuilder instance
func (fac *SignatureBuilderFactory) Create() user.SignatureBuilder {
	out := createSignatureBuilder(fac.sigBuilderFactory, fac.htBuilderFactory, fac.metaDataBuilderFactory)
	return out
}
