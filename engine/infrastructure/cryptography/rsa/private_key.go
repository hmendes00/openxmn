package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	cryptography "github.com/XMNBlockchain/exmachina-network/engine/domain/cryptography"
)

// PrivateKey represents a private key
type PrivateKey struct {
	key *rsa.PrivateKey
}

func createPrivateKey(key *rsa.PrivateKey) cryptography.PrivateKey {
	out := PrivateKey{
		key: key,
	}

	return &out
}

// GetKey returns the *rsa.Privatekey
func (pk *PrivateKey) GetKey() *rsa.PrivateKey {
	return pk.key
}

// GetPublicKey returns the PublicKey of the PrivateKey
func (pk *PrivateKey) GetPublicKey() cryptography.PublicKey {
	pubKey := createPublicKey(&pk.key.PublicKey)
	return pubKey
}

// String represents a string representation of the PrivateKey
func (pk *PrivateKey) String() string {
	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pk.key),
	}

	mem := pem.EncodeToMemory(privateKey)
	encodedStr := base64.StdEncoding.EncodeToString(mem)
	return encodedStr
}
