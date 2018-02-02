package rsa

import (
	"crypto/rsa"
	"errors"

	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
)

type publicKeyBuilder struct {
	encodedString string
	key           *rsa.PublicKey
}

func createPublicKeyBuilder() cryptography.PublicKeyBuilder {
	out := publicKeyBuilder{
		encodedString: "",
		key:           nil,
	}

	return &out
}

// Create initializes a PublicKeyBuilder instance
func (build *publicKeyBuilder) Create() cryptography.PublicKeyBuilder {
	build.encodedString = ""
	build.key = nil
	return build
}

// WithEncodedString adds an encodedString to the PublicKeyBuilder instance
func (build *publicKeyBuilder) WithEncodedString(encodedString string) cryptography.PublicKeyBuilder {
	build.encodedString = encodedString
	return build
}

// WithKey adds a *rsa.PublicKey instance to the PublicKeyBuilder instance
func (build *publicKeyBuilder) WithKey(key *rsa.PublicKey) cryptography.PublicKeyBuilder {
	build.key = key
	return build
}

// Now builds a PublicKey instance
func (build *publicKeyBuilder) Now() (cryptography.PublicKey, error) {

	if build.encodedString != "" && build.key != nil {
		return nil, errors.New("both the encodedString and the *rsa.PublicKey were set.  Only one of them can be set")
	}

	if build.encodedString == "" && build.key == nil {
		return nil, errors.New("both the encodedString and the *rsa.PublicKey are nil.  One of them is mandatory")
	}

	if build.key != nil {
		out := createPublicKey(build.key)
		return out, nil
	}

	out, outErr := createPublicKeyFromEncodedString(build.encodedString)
	return out, outErr
}
