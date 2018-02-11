package domain

// ObjectBuilderFactory represents a stored ObjectBuilderFactory
type ObjectBuilderFactory interface {
	Create() ObjectBuilder
}
