package domain

// FileRepository represents a file repository
type FileRepository interface {
	Retrieve(dirPath string, fileName string) (File, error)
	RetrieveAll(dirPath string, fileNames []string) ([]File, error)
}
