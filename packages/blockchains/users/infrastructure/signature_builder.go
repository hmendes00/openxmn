package infrastructure

import (
	"errors"
	"strconv"
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	uuid "github.com/satori/go.uuid"
)

type signatureBuilder struct {
	sigBuilderFactory      cryptography.SignatureBuilderFactory
	htBuilderFactory       hashtrees.HashTreeBuilderFactory
	metaDataBuilderFactory metadata.MetaDataBuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	v                      interface{}
	pk                     cryptography.PrivateKey
	sig                    cryptography.Signature
	usr                    users.User
	crOn                   *time.Time
}

func createSignatureBuilder(sigBuilderFactory cryptography.SignatureBuilderFactory, htBuilderFactory hashtrees.HashTreeBuilderFactory, metaDataBuilderFactory metadata.MetaDataBuilderFactory) users.SignatureBuilder {
	out := signatureBuilder{
		sigBuilderFactory:      sigBuilderFactory,
		htBuilderFactory:       htBuilderFactory,
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:   nil,
		met:  nil,
		v:    nil,
		pk:   nil,
		sig:  nil,
		usr:  nil,
		crOn: nil,
	}
	return &out
}

// Create creates a new Signature instance
func (build *signatureBuilder) Create() users.SignatureBuilder {
	build.id = nil
	build.met = nil
	build.v = nil
	build.pk = nil
	build.sig = nil
	build.usr = nil
	build.crOn = nil
	return build
}

// WithID adds an ID to the SignatureBuilder instance
func (build *signatureBuilder) WithID(id *uuid.UUID) users.SignatureBuilder {
	build.id = id
	return build
}

// WithMetaData adds MetaData to the SignatureBuilder instance
func (build *signatureBuilder) WithMetaData(met metadata.MetaData) users.SignatureBuilder {
	build.met = met
	return build
}

// WithPrivateKey adds a PrivateKey to the SignatureBuilder
func (build *signatureBuilder) WithPrivateKey(pk cryptography.PrivateKey) users.SignatureBuilder {
	build.pk = pk
	return build
}

// WithInterface adds an interface to sign to the SignatureBuilder
func (build *signatureBuilder) WithInterface(v interface{}) users.SignatureBuilder {
	build.v = v
	return build
}

// WithSignature adds a cryptographic Signature to the SignatureBuilder
func (build *signatureBuilder) WithSignature(sig cryptography.Signature) users.SignatureBuilder {
	build.sig = sig
	return build
}

// WithUser adds a User to the SignatureBuilder
func (build *signatureBuilder) WithUser(usr users.User) users.SignatureBuilder {
	build.usr = usr
	return build
}

// CreatedOn adds a creation time to the SignatureBuilder
func (build *signatureBuilder) CreatedOn(crOn time.Time) users.SignatureBuilder {
	build.crOn = &crOn
	return build
}

// Now builds a user Signature
func (build *signatureBuilder) Now() (users.Signature, error) {

	if build.usr == nil {
		return nil, errors.New("the User is mandatory in order to build a user Signature")
	}

	if build.sig == nil && build.v == nil {
		return nil, errors.New("the cryptographic signature or the instance is mandatory in order to build a user Signature")
	}

	if build.v != nil {
		pk := build.pk.GetKey()
		sig, sigErr := build.sigBuilderFactory.Create().Create().WithInterface(build.v).WithPrivateKey(pk).Now()
		if sigErr != nil {
			return nil, sigErr
		}

		build.sig = sig
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a Signature instance")
		}

		if build.crOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build a Signature instance")
		}

		blocks := [][]byte{
			build.id.Bytes(),
			[]byte(strconv.Itoa(int(build.crOn.UnixNano()))),
			build.usr.GetMetaData().GetHashTree().GetHash().Get(),
			[]byte(build.sig.String()),
		}

		ht, htErr := build.htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
		if htErr != nil {
			return nil, htErr
		}

		met, metErr := build.metaDataBuilderFactory.Create().Create().WithID(build.id).WithHashTree(ht).CreatedOn(*build.crOn).Now()
		if metErr != nil {
			return nil, metErr
		}

		build.met = met
	}

	if build.met == nil {
		return nil, errors.New("the MetaData is mandatory in order to build a Signature instance")
	}

	out := createSignature(build.met.(*concrete_metadata.MetaData), build.sig.(*concrete_cryptography.Signature), build.usr.(*User))
	return out, nil
}
