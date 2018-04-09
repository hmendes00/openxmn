package validated

// SignedBlockBuilderFactory represents a stored validated SignedBlockBuilderFactory
type SignedBlockBuilderFactory interface {
	Create() SignedBlockBuilder
}
