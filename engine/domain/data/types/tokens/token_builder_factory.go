package tokens

// TokenBuilderFactory represents a token builder factory
type TokenBuilderFactory interface {
	Create() TokenBuilder
}
