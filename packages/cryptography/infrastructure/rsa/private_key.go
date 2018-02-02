package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
)

type privateKey struct {
	key *rsa.PrivateKey
}

func createPrivateKey(key *rsa.PrivateKey) cryptography.PrivateKey {
	out := privateKey{
		key: key,
	}

	return &out
}

// GetKey returns the *rsa.Privatekey
func (pk *privateKey) GetKey() *rsa.PrivateKey {
	return pk.key
}

// String represents a string representation of the PrivateKey
func (pk *privateKey) String() string {
	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pk.key),
	}

	mem := pem.EncodeToMemory(privateKey)
	encodedStr := base64.StdEncoding.EncodeToString(mem)
	return encodedStr
}
