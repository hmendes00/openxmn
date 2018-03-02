package domain

// ChainBuilderFactory represents a chain builder factory
type ChainBuilderFactory interface {
	Create() ChainBuilder
}
