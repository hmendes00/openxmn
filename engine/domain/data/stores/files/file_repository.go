package files

// FileRepository represents a stored file repository
type FileRepository interface {
	Retrieve(filePath string) (File, error)
	RetrieveAll(dirPath string) ([]File, error)
}
