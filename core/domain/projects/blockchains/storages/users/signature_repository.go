package users

// SignatureRepository represents a stored signature repository
type SignatureRepository interface {
	Retrieve(dirPath string) (Signature, error)
	RetrieveAll(dirPath string) ([]Signature, error)
}
