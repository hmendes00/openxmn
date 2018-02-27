package domain

// SignatureBuilderFactory represents a SignatureBuilder factory
type SignatureBuilderFactory interface {
	Create() SignatureBuilder
}
