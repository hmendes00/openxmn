package users

import (
	cryptography "github.com/XMNBlockchain/exmachina-network/core/domain/cryptography"
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/hashtrees"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	user "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/users"
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
