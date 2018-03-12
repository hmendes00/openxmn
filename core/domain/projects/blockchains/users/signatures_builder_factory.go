package domain

// SignaturesBuilderFactory represents a SignaturesBulderFactory
type SignaturesBuilderFactory interface {
	Create() SignaturesBuilder
}
