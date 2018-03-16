package users

import (
	"errors"

	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/users"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/files"
)

type signaturesBuilder struct {
	met  stored_files.File
	sigs []stored_users.Signature
}

func createSignaturesBuilder() stored_users.SignaturesBuilder {
	out := signaturesBuilder{
		met:  nil,
		sigs: nil,
	}

	return &out
}

// Create initializes the SignaturesBuilder
func (build *signaturesBuilder) Create() stored_users.SignaturesBuilder {
	build.met = nil
	build.sigs = nil
	return build
}

// WithMetaData adds MetaData to the SignaturesBuilder
func (build *signaturesBuilder) WithMetaData(met stored_files.File) stored_users.SignaturesBuilder {
	build.met = met
	return build
}

// WithSignatures adds []Signature to the SignaturesBuilder
func (build *signaturesBuilder) WithSignatures(sigs []stored_users.Signature) stored_users.SignaturesBuilder {
	build.sigs = sigs
	return build
}

// Now builds a new Signatures instance
func (build *signaturesBuilder) Now() (stored_users.Signatures, error) {
	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a Signatures instance")
	}

	if build.sigs == nil {
		return nil, errors.New("the []Signature is mandatory in order to build a Signatures instance")
	}

	sigs := []*Signature{}
	for _, oneSig := range build.sigs {
		sigs = append(sigs, oneSig.(*Signature))
	}

	out := createSignatures(build.met.(*concrete_stored_files.File), sigs)
	return out, nil
}
