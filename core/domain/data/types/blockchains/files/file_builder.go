package domain

// FileBuilder represents a stored file builder
type FileBuilder interface {
	Create() FileBuilder
	WithData(data []byte) FileBuilder
	WithDirPath(dirPath string) FileBuilder
	WithFileName(fileName string) FileBuilder
	WithExtension(ext string) FileBuilder
	Now() (File, error)
}
