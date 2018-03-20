package users

// SignatureBuilderFactory represents a stored signature builder factory
type SignatureBuilderFactory interface {
	Create() SignatureBuilder
}
