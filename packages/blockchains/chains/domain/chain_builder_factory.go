package domain

// ChainBuilderFactory represents the blockchain builder factory
type ChainBuilderFactory interface {
	Create() ChainBuilder
}
