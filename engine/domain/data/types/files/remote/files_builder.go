package remote

// FilesBuilder represents a remote files builder
type FilesBuilder interface {
	Create() FilesBuilder
	WithFiles(fils []File) FilesBuilder
	Now() (Files, error)
}
