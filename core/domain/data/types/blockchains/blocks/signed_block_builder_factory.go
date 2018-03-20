package domain

// SignedBlockBuilderFactory represents a SignedBlockBuilder factory
type SignedBlockBuilderFactory interface {
	Create() SignedBlockBuilder
}
