package domain

import "crypto/rsa"

// PublicKeyBuilder represents a PublicKey builder
type PublicKeyBuilder interface {
	Create() PublicKeyBuilder
	WithEncodedString(encodedString string) PublicKeyBuilder
	WithKey(key *rsa.PublicKey) PublicKeyBuilder
	Now() (PublicKey, error)
}
