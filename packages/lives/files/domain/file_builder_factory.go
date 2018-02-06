package domain

// FileBuilderFactory represents a stored file builder factory
type FileBuilderFactory interface {
	Create() FileBuilder
}
