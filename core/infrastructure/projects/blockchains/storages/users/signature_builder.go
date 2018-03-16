package users

import (
	"errors"

	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/users"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/files"
)

type signatureBuilder struct {
	met stored_files.File
	sig stored_files.File
	usr stored_users.User
}

func createSignatureBuilder() stored_users.SignatureBuilder {
	out := signatureBuilder{
		met: nil,
		sig: nil,
		usr: nil,
	}

	return &out
}

// Create initializes the SignatureBuilder
func (build *signatureBuilder) Create() stored_users.SignatureBuilder {
	build.met = nil
	build.sig = nil
	build.usr = nil
	return build
}

// WithMetaData adds MetaData to the SignatureBuilder
func (build *signatureBuilder) WithMetaData(met stored_files.File) stored_users.SignatureBuilder {
	build.met = met
	return build
}

// WithSignature adds a Signature to the SignatureBuilder
func (build *signatureBuilder) WithSignature(sig stored_files.File) stored_users.SignatureBuilder {
	build.sig = sig
	return build
}

// WithUser adds a User to the SignatureBuilder
func (build *signatureBuilder) WithUser(usr stored_users.User) stored_users.SignatureBuilder {
	build.usr = usr
	return build
}

// Now builds a new Signature instance
func (build *signatureBuilder) Now() (stored_users.Signature, error) {
	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a Signature instance")
	}

	if build.sig == nil {
		return nil, errors.New("the Signature is mandatory in order to build a Signature instance")
	}

	if build.usr == nil {
		return nil, errors.New("the User is mandatory in order to build a Signature instance")
	}

	out := createSignature(build.met.(*concrete_stored_files.File), build.sig.(*concrete_stored_files.File), build.usr.(*User))
	return out, nil
}
