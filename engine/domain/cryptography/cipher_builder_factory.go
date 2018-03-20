package cryptography

// CipherBuilderFactory represents a cipher builder factory
type CipherBuilderFactory interface {
	Create() CipherBuilder
}
