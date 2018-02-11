package domain

// TreeBuilderFactory represents a TreeBuilderFactory
type TreeBuilderFactory interface {
	Create() TreeBuilder
}
