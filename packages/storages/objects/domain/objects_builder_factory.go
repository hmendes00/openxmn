package domain

// ObjectsBuilderFactory represents an ObjectsBuilder factory
type ObjectsBuilderFactory interface {
	Create() ObjectsBuilder
}
