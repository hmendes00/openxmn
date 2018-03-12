package domain

// SignaturesRepository represents a signatures repository
type SignaturesRepository interface {
	Retrieve(dirPath string) (Signatures, error)
}
