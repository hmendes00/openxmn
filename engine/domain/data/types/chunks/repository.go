package chunks

// Repository represents a Chunks repository
type Repository interface {
	Retrieve(dirPath string) (Chunks, error)
	RetrieveAll(dirPath string) ([]Chunks, error)
}
