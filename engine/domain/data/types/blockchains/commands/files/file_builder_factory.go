package files

// FileBuilderFactory represents a file builder factory
type FileBuilderFactory interface {
	Create() FileBuilder
}
