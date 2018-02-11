package domain

// ObjectBuilderFactory represents an ObjectBuilderFactory
type ObjectBuilderFactory interface {
	Create() ObjectBuilder
}
