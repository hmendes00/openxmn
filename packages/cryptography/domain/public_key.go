package domain

import "crypto/rsa"

// PublicKey represents a public key
type PublicKey interface {
	String() (string, error)
	Compare(pk PublicKey) bool
	GetKey() *rsa.PublicKey
}
