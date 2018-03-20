package remote

// FileRepository represents a remote file repository
type FileRepository interface {
	Retrieve(filePath string) (File, error)
	RetrieveAll(dirPath string) ([]File, error)
}
