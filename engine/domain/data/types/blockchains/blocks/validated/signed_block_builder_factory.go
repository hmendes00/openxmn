package domain

// SignedBlockBuilderFactory represents a signed block builder factory
type SignedBlockBuilderFactory interface {
	Create() SignedBlockBuilder
}
