package domain

// SignedBlockBuilderFactory represents a stored signed block builder factory
type SignedBlockBuilderFactory interface {
	Create() SignedBlockBuilder
}
