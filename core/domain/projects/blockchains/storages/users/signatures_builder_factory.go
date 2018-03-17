package users

// SignaturesBuilderFactory represents a SignaturesBuilder factory
type SignaturesBuilderFactory interface {
	Create() SignaturesBuilder
}
