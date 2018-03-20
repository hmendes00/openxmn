package domain

// HashTreeBuilderFactory represents a HashTreeBuilder factory
type HashTreeBuilderFactory interface {
	Create() HashTreeBuilder
}
