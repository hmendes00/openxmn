package domain

// TreesBuilderFactory represents a TreesBuilder factory
type TreesBuilderFactory interface {
	Create() TreesBuilder
}
