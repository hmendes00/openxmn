package remote

// FilesBuilderFactory represents a remote files builder factory
type FilesBuilderFactory interface {
	Create() FilesBuilder
}
