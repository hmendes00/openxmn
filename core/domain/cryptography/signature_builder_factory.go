package cryptography

// SignatureBuilderFactory represents a SignatureBuilder
type SignatureBuilderFactory interface {
	Create() SignatureBuilder
}
