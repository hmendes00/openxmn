package cryptography

// PrivateKeyBuilderFactory represents a private key builder factory
type PrivateKeyBuilderFactory interface {
	Create() PrivateKeyBuilder
}
