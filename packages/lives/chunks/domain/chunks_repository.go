package domain

// ChunksRepository represents a Chunks repository
type ChunksRepository interface {
	Retrieve(dirPath string, h string) (Chunks, error)
	RetrieveAll(dirPath []string, h []string) ([]Chunks, error)
}
