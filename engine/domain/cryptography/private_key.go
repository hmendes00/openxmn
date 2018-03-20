package cryptography

import (
	"crypto/rsa"
)

// PrivateKey represents a private key
type PrivateKey interface {
	GetKey() *rsa.PrivateKey
	GetPublicKey() PublicKey
	String() string
}
