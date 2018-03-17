package users

// SignaturesRepository represents a stored signatures repository
type SignaturesRepository interface {
	Retrieve(dirPath string) (Signatures, error)
}
