package domain

import "crypto/rsa"

// PrivateKey represents a private key
type PrivateKey interface {
	GetKey() *rsa.PrivateKey
	String() string
}
