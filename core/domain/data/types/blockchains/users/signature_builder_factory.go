package domain

// SignatureBuilderFactory represents the SignatureBuilder factory
type SignatureBuilderFactory interface {
	Create() SignatureBuilder
}
