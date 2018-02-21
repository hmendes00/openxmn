package infrastructure

import (
	"errors"

	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
	uuid "github.com/satori/go.uuid"
)

type signatureBuilder struct {
	sigBuilderFactory  cryptography.SignatureBuilderFactory
	userBuilderFactory users.UserBuilderFactory
	v                  interface{}
	pk                 cryptography.PrivateKey
	sig                cryptography.Signature
	encodedSig         string
	userID             *uuid.UUID
}

func createSignatureBuilder(sigBuilderFactory cryptography.SignatureBuilderFactory, userBuilderFactory users.UserBuilderFactory) users.SignatureBuilder {
	out := signatureBuilder{
		sigBuilderFactory:  sigBuilderFactory,
		userBuilderFactory: userBuilderFactory,
		v:                  nil,
		pk:                 nil,
		sig:                nil,
		encodedSig:         "",
		userID:             nil,
	}
	return &out
}

// Create creates a new Signature instance
func (build *signatureBuilder) Create() users.SignatureBuilder {
	build.userID = nil
	build.v = nil
	build.pk = nil
	build.sig = nil
	build.encodedSig = ""
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

// WithEncodedSignature adds a cryptographic Signature to the SignatureBuilder
func (build *signatureBuilder) WithEncodedSignature(encodedSig string) users.SignatureBuilder {
	build.encodedSig = encodedSig
	return build
}

// WithSignature adds a cryptographic Signature to the SignatureBuilder
func (build *signatureBuilder) WithSignature(sig cryptography.Signature) users.SignatureBuilder {
	build.sig = sig
	return build
}

// WithUserID adds a PointerID to the SignatureBuilder
func (build *signatureBuilder) WithUserID(userID *uuid.UUID) users.SignatureBuilder {
	build.userID = userID
	return build
}

// Now builds a user Signature
func (build *signatureBuilder) Now() (users.Signature, error) {

	if build.encodedSig != "" {
		sig, sigErr := build.sigBuilderFactory.Create().Create().WithEncodedSignature(build.encodedSig).Now()
		if sigErr != nil {
			return nil, sigErr
		}

		build.sig = sig
	}

	if build.userID == nil {
		return nil, errors.New("the userID is mandatory in order to build a user Signature")
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

	pubKey := build.sig.GetPublicKey()
	usr, usrErr := build.userBuilderFactory.Create().Create().WithID(*build.userID).WithPublicKey(pubKey).Now()
	if usrErr != nil {
		return nil, usrErr
	}

	out := createSignature(build.sig.(*concrete_cryptography.Signature), usr.(*User))
	return out, nil
}
