package remote

// FileBuilderFactory represents a remote file builder factory
type FileBuilderFactory interface {
	Create() FileBuilder
}
