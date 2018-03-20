package files

// FileBuilderFactory represents a FileBuilder factory
type FileBuilderFactory interface {
	Create() FileBuilder
}
