package domain

// Signature represents a signature
type Signature interface {
	String() string
	GetPublicKey() PublicKey
}
