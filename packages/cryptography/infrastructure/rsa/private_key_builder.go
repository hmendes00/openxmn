package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"

	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
)

type privateKeyBuilder struct {
	encodedStr string
	pk         *rsa.PrivateKey
}

func createPrivateKeyBuilder() cryptography.PrivateKeyBuilder {
	out := privateKeyBuilder{
		encodedStr: "",
		pk:         nil,
	}

	return &out
}

// Create initializes the PrivateKeyBuilder instance
func (build *privateKeyBuilder) Create() cryptography.PrivateKeyBuilder {
	build.encodedStr = ""
	build.pk = nil
	return build
}

// WithEncodedString adds an encodedString to the PrivateKeyBuilder instance
func (build *privateKeyBuilder) WithEncodedString(encodedStr string) cryptography.PrivateKeyBuilder {
	build.encodedStr = encodedStr
	return build
}

// WithKey adds a *rsa.PrivateKey instance to the PrivateKeyBuilder instance
func (build *privateKeyBuilder) WithKey(pk *rsa.PrivateKey) cryptography.PrivateKeyBuilder {
	build.pk = pk
	return build
}

// Now builds a PrivateKey instance
func (build *privateKeyBuilder) Now() (cryptography.PrivateKey, error) {

	if build.encodedStr != "" && build.pk != nil {
		return nil, errors.New("both the encodedString and the *rsa.PrivateKey were set.  Only one of them can be set")
	}

	if build.encodedStr == "" && build.pk == nil {
		bitSize := 4096
		reader := rand.Reader

		pk, pkErr := rsa.GenerateKey(reader, bitSize)
		if pkErr != nil {
			return nil, pkErr
		}

		out := createPrivateKey(pk)
		return out, nil
	}

	if build.encodedStr != "" {

		privAsBytes, privAsBytesErr := base64.StdEncoding.DecodeString(string(build.encodedStr))
		if privAsBytesErr != nil {
			return nil, privAsBytesErr
		}

		priv, privErr := func(pkey []byte) (*rsa.PrivateKey, error) {
			block, remaining := pem.Decode(pkey)
			if block == nil {
				return x509.ParsePKCS1PrivateKey(remaining)
			}

			return x509.ParsePKCS1PrivateKey(block.Bytes)
		}(privAsBytes)

		if privErr != nil {
			return nil, privErr
		}

		out := createPrivateKey(priv)
		return out, nil
	}

	out := createPrivateKey(build.pk)
	return out, nil
}
