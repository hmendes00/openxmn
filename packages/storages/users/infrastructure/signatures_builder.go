package infrastructure

import (
	"errors"

	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
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

	out := createSignatures(build.met, build.sigs)
	return out, nil
}
