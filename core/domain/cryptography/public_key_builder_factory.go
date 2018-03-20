package cryptography

// PublicKeyBuilderFactory represents a PublicKeyBuilder factory
type PublicKeyBuilderFactory interface {
	Create() PublicKeyBuilder
}
