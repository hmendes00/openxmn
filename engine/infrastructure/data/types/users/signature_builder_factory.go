package users

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	user "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
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
