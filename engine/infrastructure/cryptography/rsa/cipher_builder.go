package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"errors"
	"hash"

	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
)

type cipherBuilder struct {
	sigBuilderFactory cryptography.SignatureBuilderFactory
	pk                cryptography.PrivateKey
	pubKey            cryptography.PublicKey
	txt               []byte
	hash              hash.Hash
	cipherTxt         []byte
	label             []byte
	sig               cryptography.Signature
}

func createCipherBuilder(sigBuilderFactory cryptography.SignatureBuilderFactory, pk cryptography.PrivateKey) cryptography.CipherBuilder {
	out := cipherBuilder{
		sigBuilderFactory: sigBuilderFactory,
		pk:                pk,
		pubKey:            nil,
		txt:               nil,
		hash:              nil,
		cipherTxt:         nil,
		label:             nil,
		sig:               nil,
	}

	return &out
}

// Create initializes the CipherBuilder
func (build *cipherBuilder) Create() cryptography.CipherBuilder {
	build.pubKey = nil
	build.txt = nil
	build.hash = nil
	build.cipherTxt = nil
	build.label = nil
	build.sig = nil
	return build
}

// WithPublicKey adds a PublicKey to the CipherBuilder
func (build *cipherBuilder) WithPublicKey(pubKey cryptography.PublicKey) cryptography.CipherBuilder {
	build.pubKey = pubKey
	return build
}

// WithText adds a text to the CipherBuilder
func (build *cipherBuilder) WithText(txt []byte) cryptography.CipherBuilder {
	build.txt = txt
	return build
}

// WithHash adds a hash to the CipherBuilder
func (build *cipherBuilder) WithHash(h hash.Hash) cryptography.CipherBuilder {
	build.hash = h
	return build
}

// WithCipherText adds a cipher text to the CipherBuilder
func (build *cipherBuilder) WithCipherText(cipherText []byte) cryptography.CipherBuilder {
	build.cipherTxt = cipherText
	return build
}

// WithLabel adds a label to the CipherBuilder
func (build *cipherBuilder) WithLabel(label []byte) cryptography.CipherBuilder {
	build.label = label
	return build
}

// WithSignature adds a signature to the CipherBuilder
func (build *cipherBuilder) WithSignature(sig cryptography.Signature) cryptography.CipherBuilder {
	build.sig = sig
	return build
}

// Now builds a cipher instance
func (build *cipherBuilder) Now() (cryptography.Cipher, error) {
	if build.txt != nil && build.pubKey != nil && build.label != nil {

		hash := sha256.New()
		cipherText, cipherTextErr := rsa.EncryptOAEP(hash, rand.Reader, build.pubKey.GetKey(), build.txt, build.label)
		if cipherTextErr != nil {
			return nil, cipherTextErr
		}

		sig, sigErr := build.sigBuilderFactory.Create().Create().WithPrivateKey(build.pk.GetKey()).WithData(cipherText).Now()
		if sigErr != nil {
			return nil, sigErr
		}

		build.hash = hash
		build.sig = sig
		build.cipherTxt = cipherText
	}

	if build.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Cipher instance")
	}

	if build.cipherTxt == nil {
		return nil, errors.New("the cipherText is mandatory in order to build a Cipher instance")
	}

	if build.label == nil {
		return nil, errors.New("the label is mandatory in order to build a Cipher instance")
	}

	if build.sig == nil {
		return nil, errors.New("the signature is mandatory in order to build a Cipher instance")
	}

	cipher := createCipher(build.hash, build.cipherTxt, build.label, build.sig.(*Signature))
	return cipher, nil
}
