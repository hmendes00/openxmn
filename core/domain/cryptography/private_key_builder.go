package cryptography

import "crypto/rsa"

// PrivateKeyBuilder represents a private key builder
type PrivateKeyBuilder interface {
	Create() PrivateKeyBuilder
	WithEncodedString(encodedString string) PrivateKeyBuilder
	WithKey(pk *rsa.PrivateKey) PrivateKeyBuilder
	Now() (PrivateKey, error)
}
