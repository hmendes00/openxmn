package domain

// ChunksRepository represents a Chunks repository
type ChunksRepository interface {
	Retrieve(dirPath string) (Chunks, error)
}
