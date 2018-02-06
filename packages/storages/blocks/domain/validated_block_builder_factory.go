package domain

// ValidatedBlockBuilderFactory represents a stored validated block builder factory
type ValidatedBlockBuilderFactory interface {
	Create() ValidatedBlockBuilder
}
