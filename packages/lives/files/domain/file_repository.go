package domain

// FileRepository represents a file repository
type FileRepository interface {
	Retrieve(dirPath string, h string) (File, error)
	RetrieveAll(dirPath string, h []string) ([]File, error)
}
